package UtilNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const RotorFactoryName = "Node.UtilNode.RotorNode"

func init() {
	Scene.NodeFactory.AddType(RotorFactoryName, reflect.TypeOf((*RotorNode)(nil)).Elem())
}

type RotorNode struct {
	Scene.INode
	Speed float32
	Axis  GeometryMath.Vector3
}

func (node *RotorNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetLocalTransformation(node.GetLocalTransformation().Mul(GeometryMath.RotateY(node.Speed * timeDelta)))

	return err
}

func (node *RotorNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *RotorNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	yamlConfig := struct {
		Speed float32              `yaml:"speed"`
		Axis  GeometryMath.Vector3 `yaml:"axis"`
	}{
		Speed: node.Speed,
		Axis:  node.Axis,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Speed = yamlConfig.Speed
	node.Axis = yamlConfig.Axis

	if node.Axis == (GeometryMath.Vector3{}) {
		node.Axis = GeometryMath.Vector3{0, 1, 0}
	}

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
