package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light/SpotLight"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const SpotLightNodeFactoryName = "Node.LightNode.SpotLightNode"
const SpotLightUBOFactoryName = "spotLight"

func init() {
	YamlFactory.NodeFactory[SpotLightNodeFactoryName] = reflect.TypeOf((*SpotLightNodeConfig)(nil)).Elem()
	UniformBufferFactory.AddType(SpotLightUBOFactoryName, reflect.TypeOf((*SpotLight.UniformBuffer)(nil)).Elem())
}

type SpotLightNodeConfig struct {
	Scene.NodeConfig
	SpotLight.SpotLight `yaml:"spotLight"`
	UBO                 string `yaml:"uniformBuffer"`
}

func (config *SpotLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	ubo, err := UniformBufferFactory.Get(config.UBO)
	if err != nil {
		return nil, err
	}

	lightUbo := ubo.(SpotLight.IUniformBuffer)
	light, err := lightUbo.GetNewElement()
	if err != nil {
		return nil, err
	}

	light.Set(config.SpotLight)

	node := &SpotLightNode{
		INode:      nodeBase,
		ISpotLight: light,
		Config:     config,
	}

	return node, nil
}

type SpotLightNode struct {
	Scene.INode
	SpotLight.ISpotLight

	Config *SpotLightNodeConfig
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetPosition(*node.GetGlobalPosition())
	node.SetDirection(*node.GetGlobalTransformation().Inverse().Transpose().MulVector(&node.Config.Direction).Normalize())

	return err
}
