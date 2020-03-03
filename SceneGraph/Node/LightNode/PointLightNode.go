package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light/PointLight"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const PointLightNodeFactoryName = "Node.LightNode.PointLightNode"
const PointLightUBOFactoryName = "pointLight"

func init() {
	NodeFactory.AddType(PointLightNodeFactoryName, reflect.TypeOf((*PointLightNodeConfig)(nil)).Elem())
	UniformBufferFactory.AddType(PointLightUBOFactoryName, reflect.TypeOf((*PointLight.UniformBuffer)(nil)).Elem())
}

type PointLightNodeConfig struct {
	Scene.NodeConfig
	PointLight.PointLight `yaml:"pointLight"`
	UBO                   string `yaml:"uniformBuffer"`
}

func (config *PointLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	ubo, err := UniformBufferFactory.Get(config.UBO)
	if err != nil {
		return nil, err
	}

	lightUbo := ubo.(PointLight.IUniformBuffer)
	light, err := lightUbo.GetNewElement()
	if err != nil {
		return nil, err
	}

	light.Set(config.PointLight)

	node := &PointLightNode{
		INode:       nodeBase,
		IPointLight: light,
		Config:      config,
	}

	return node, nil
}

type PointLightNode struct {
	Scene.INode
	PointLight.IPointLight
	Config *PointLightNodeConfig
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.IPointLight.SetPosition(*node.GetGlobalPosition())

	return err
}
