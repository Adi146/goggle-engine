package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const (
	DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
)

func init() {
	Scene.NodeFactory.AddType(DirectionalLightNodeFactoryName, reflect.TypeOf((*DirectionalLightNode)(nil)).Elem())
}

type DirectionalLightNode struct {
	Scene.INode
	DirectionalLight Light.IDirectionalLight
	InitDirection    GeometryMath.Vector3
	ShadowMap        struct {
		Camera      Camera.ICamera
		Shader      Shader.IShaderProgram
		FrameBuffer ShadowMapShader.ShadowMapBuffer
	}
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	newDirection := *node.GetGlobalTransformation().MulVector(&node.InitDirection).Normalize()

	node.DirectionalLight.SetDirection(newDirection)
	node.ShadowMap.Camera.SetViewMatrix(*GeometryMath.LookAt(newDirection.Invert(), &GeometryMath.Vector3{0, 0, 0}, &GeometryMath.Vector3{0, 1, 0}))

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *DirectionalLightNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if invoker == node {
		return nil
	}
	defer FrameBuffer.GetCurrentFrameBuffer().Bind()
	defer Function.GetCurrentCullFunction().Set()
	defer Function.GetCurrentDepthFunction().Set()
	defer Function.GetCurrentBlendFunction().Set()

	node.ShadowMap.FrameBuffer.Bind()
	Function.Front.Set()
	Function.Less.Set()
	Function.DisabledBlend.Set()

	node.ShadowMap.FrameBuffer.Clear()

	if shader != nil {
		defer shader.Bind()
	}

	return scene.Draw(node.ShadowMap.Shader, node, scene)
}

func (node *DirectionalLightNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	if node.DirectionalLight == nil {
		light := Light.UBODirectionalLight{}

		node.DirectionalLight = &light
		node.ShadowMap.Camera = &light.CameraSection
	}
	if err := value.Decode(node.DirectionalLight); err != nil {
		return err
	}

	if node.InitDirection == (GeometryMath.Vector3{}) {
		node.InitDirection = node.DirectionalLight.GetDirection()
	}

	type shadowMapConfig struct {
		Shader      Shader.Ptr                      `yaml:"shader"`
		FrameBuffer ShadowMapShader.ShadowMapBuffer `yaml:"frameBuffer"`
	}
	yamlConfig := struct {
		ShadowMap shadowMapConfig `yaml:"shadowMap"`
	}{
		ShadowMap: shadowMapConfig{
			Shader: Shader.Ptr{
				IShaderProgram: node.ShadowMap.Shader,
			},
			FrameBuffer: node.ShadowMap.FrameBuffer,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.ShadowMap.Shader = yamlConfig.ShadowMap.Shader
	node.ShadowMap.FrameBuffer = yamlConfig.ShadowMap.FrameBuffer

	return nil
}
