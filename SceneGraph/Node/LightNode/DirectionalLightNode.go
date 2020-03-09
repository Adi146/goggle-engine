package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/MatrixFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
const DirectionalLightUBOFactoryName = "directionalLight"
const ShadowMapShaderFactoryName = "shadowMapShader"
const ShadowMapFramebufferName = "shadowMapBuffer"

func init() {
	NodeFactory.AddType(DirectionalLightNodeFactoryName, reflect.TypeOf((*DirectionalLightNodeConfig)(nil)).Elem())
	UniformBufferFactory.AddType(DirectionalLightUBOFactoryName, reflect.TypeOf((*DirectionalLight.UniformBuffer)(nil)).Elem())
	ShaderFactory.AddType(ShadowMapShaderFactoryName, ShadowMapShader.NewIShaderProgram)
	FrameBufferFactory.AddType(ShadowMapFramebufferName, reflect.TypeOf((*ShadowMapShader.ShadowMapBuffer)(nil)).Elem())
}

type DirectionalLightNodeConfig struct {
	Scene.NodeConfig
	DirectionalLight.DirectionalLight `yaml:"directionalLight"`
	UboConfig                         UniformBufferFactory.Config `yaml:"uniformBuffer"`
	ShadowMap                         struct {
		ProjectionMatrix MatrixFactory.Config `yaml:"projection"`
	} `yaml:"shadowMap"`
}

func (config *DirectionalLightNodeConfig) Create() (Scene.INode, error) {
	config.SetDefaults()

	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	light := config.UboConfig.IUniformBuffer.(DirectionalLight.IDirectionalLight)
	config.DirectionalLight.ProjectionMatrix = config.ShadowMap.ProjectionMatrix.Matrix4x4
	light.Set(config.DirectionalLight)

	node := &DirectionalLightNode{
		INode:             nodeBase,
		IDirectionalLight: light,
		Config:            config,
	}

	return node, nil
}

func (config *DirectionalLightNodeConfig) SetDefaults() {
	if config.Direction.Length() == 0 {
		config.Direction = Vector.Vector3{0, 0, -1}
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
	node.SetViewMatrix(*Matrix.LookAt(newDirection.Invert(), &Vector.Vector3{0, 0, 0}, &Vector.Vector3{0, 1, 0}))

	return err
}
