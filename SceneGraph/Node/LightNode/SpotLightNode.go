package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const SpotLightNodeFactoryName = "Node.LightNode.SpotLightNode"

func init() {
	Scene.NodeFactory.AddType(SpotLightNodeFactoryName, reflect.TypeOf((*SpotLightNode)(nil)).Elem())
}

type SpotLightNode struct {
	Scene.INode
	SpotLight     Light.UBOSpotLight
	InitDirection GeometryMath.Vector3
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SpotLight.Position.Set(*node.GetGlobalPosition())
	node.SpotLight.Direction.Set(*node.GetGlobalTransformation().Inverse().Transpose().MulVector(&node.InitDirection).Normalize())

	return err
}

func (node *SpotLightNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	if err := value.Decode(&node.SpotLight); err != nil {
		return err
	}

	if node.InitDirection == (GeometryMath.Vector3{}) {
		node.InitDirection = node.SpotLight.Direction.Get()
	}

	return nil
}
