package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const PointLightNodeFactoryName = "Node.LightNode.PointLightNode"

func init() {
	NodeFactory.AddType(PointLightNodeFactoryName, reflect.TypeOf((*PointLightNodeConfig)(nil)).Elem())
}

type PointLightNodeConfig struct {
	Scene.NodeConfig
	Light.PointLight `yaml:"pointLight"`
	UBOPointLight    Light.UBOPointLight `yaml:",inline"`
}

func (config *PointLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &PointLightNode{
		INode:       nodeBase,
		IPointLight: &config.UBOPointLight,
		Config:      config,
	}

	return node, nil
}

type PointLightNode struct {
	Scene.INode
	Light.IPointLight
	Config *PointLightNodeConfig
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.IPointLight.SetPosition(*node.GetGlobalPosition())

	return err
}
