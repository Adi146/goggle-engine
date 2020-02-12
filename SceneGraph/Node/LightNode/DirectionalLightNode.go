package LightNode

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
const DirectionalLightUBOFactoryName = "directionalLight"

func init() {
	YamlFactory.NodeFactory[DirectionalLightNodeFactoryName] = reflect.TypeOf((*DirectionalLightNodeConfig)(nil)).Elem()
	UniformBufferFactory.AddType(DirectionalLightUBOFactoryName, reflect.TypeOf((*DirectionalLight.UniformBuffer)(nil)).Elem())
}

type DirectionalLightNodeConfig struct {
	Scene.NodeConfig
	DirectionalLight.DirectionalLight `yaml:"directionalLight"`
	UBO                               string `yaml:"uniformBuffer"`
}

func (config *DirectionalLightNodeConfig) Create() (Scene.INode, error) {
	config.SetDefaults()

	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	ubo, err := UniformBufferFactory.Get(config.UBO)
	if err != nil {
		return nil, err
	}

	light := ubo.(DirectionalLight.IDirectionalLight)
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

	node.SetDirection(*node.GetGlobalTransformation().MulVector(&node.Config.Direction).Normalize())

	return err
}
