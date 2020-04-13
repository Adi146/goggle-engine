package SceneGraph

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const RotorFactoryName = "SceneGraph.Rotor"

func init() {
	Scene.NodeFactory.AddType(RotorFactoryName, reflect.TypeOf((*Rotor)(nil)).Elem())
}

type Rotor struct {
	Scene.INode
	Speed float32
}

func (node *Rotor) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetLocalTransformation(node.GetLocalTransformation().Mul(GeometryMath.RotateY(node.Speed * timeDelta)))

	return err
}

func (node *Rotor) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *Rotor) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	yamlConfig := struct {
		Speed float32 `yaml:"speed"`
	}{
		Speed: node.Speed,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Speed = yamlConfig.Speed

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
