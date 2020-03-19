package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const (
	DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
	ShadowMapFramebufferName        = "shadowMapBuffer"
)

func init() {
	Scene.NodeFactory.AddType(DirectionalLightNodeFactoryName, reflect.TypeOf((*DirectionalLightNode)(nil)).Elem())
	FrameBufferFactory.AddType(ShadowMapFramebufferName, reflect.TypeOf((*ShadowMapShader.ShadowMapBuffer)(nil)).Elem())
}

type DirectionalLightNode struct {
	Scene.INode
	DirectionalLight Light.IDirectionalLight
	InitDirection    GeometryMath.Vector3
	ShadowMap        struct {
		Camera      Camera.ICamera
		Shader      Shader.IShaderProgram
		FrameBuffer FrameBufferFactory.Config
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

	node.ShadowMap.FrameBuffer.Bind()
	defer node.ShadowMap.FrameBuffer.Unbind()
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
		Shader      Shader.Ptr                `yaml:"shader"`
		FrameBuffer FrameBufferFactory.Config `yaml:"frameBuffer"`
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
