package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const (
	DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
	ShadowMapShaderFactoryName      = "shadowMapShader"
	ShadowMapFramebufferName        = "shadowMapBuffer"
)

func init() {
	NodeFactory.AddType(DirectionalLightNodeFactoryName, reflect.TypeOf((*DirectionalLightNodeConfig)(nil)).Elem())
	ShaderFactory.AddType(ShadowMapShaderFactoryName, ShadowMapShader.NewIShaderProgram)
	FrameBufferFactory.AddType(ShadowMapFramebufferName, reflect.TypeOf((*ShadowMapShader.ShadowMapBuffer)(nil)).Elem())
}

type DirectionalLightNodeConfig struct {
	Scene.NodeConfig
	UBOSection                        DirectionalLight.UBOSection `yaml:",inline"`
	DirectionalLight.DirectionalLight `yaml:"directionalLight"`
	ShadowMap                         struct {
		Shader      ShaderFactory.Config      `yaml:"shader"`
		FrameBuffer FrameBufferFactory.Config `yaml:"frameBuffer"`
	} `yaml:"shadowMap"`
}

func (config *DirectionalLightNodeConfig) Create() (Scene.INode, error) {
	config.SetDefaults()

	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &DirectionalLightNode{
		INode:             nodeBase,
		IDirectionalLight: &config.UBOSection,
		Config:            config,
	}

	return node, nil
}

func (config *DirectionalLightNodeConfig) SetDefaults() {
	if config.Direction.Length() == 0 {
		config.Direction = GeometryMath.Vector3{0, 0, -1}
	}
}

type DirectionalLightNode struct {
	Scene.INode
	DirectionalLight.IDirectionalLight
	Config *DirectionalLightNodeConfig
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	newDirection := *node.GetGlobalTransformation().MulVector(&node.Config.Direction).Normalize()

	node.SetDirection(newDirection)
	node.SetViewMatrix(*GeometryMath.LookAt(newDirection.Invert(), &GeometryMath.Vector3{0, 0, 0}, &GeometryMath.Vector3{0, 1, 0}))

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *DirectionalLightNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if invoker == node {
		return nil
	}

	node.Config.ShadowMap.FrameBuffer.Bind()
	defer node.Config.ShadowMap.FrameBuffer.Unbind()
	node.Config.ShadowMap.FrameBuffer.Clear()

	if shader != nil {
		defer shader.Bind()
	}

	return scene.Draw(node.Config.ShadowMap.Shader, node, scene)
}
