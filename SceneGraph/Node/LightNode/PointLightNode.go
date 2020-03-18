package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const PointLightNodeFactoryName = "Node.LightNode.PointLightNode"

func init() {
	Scene.NodeFactory.AddType(PointLightNodeFactoryName, reflect.TypeOf((*PointLightNode)(nil)).Elem())
}

type PointLightNode struct {
	Scene.INode
	PointLight Light.IPointLight
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.PointLight.SetPosition(*node.GetGlobalPosition())

	return err
}

func (node *PointLightNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	if node.PointLight == nil {
		node.PointLight = &Light.UBOPointLight{}
	}
	if err := value.Decode(node.PointLight); err != nil {
		return err
	}

	return nil
}
