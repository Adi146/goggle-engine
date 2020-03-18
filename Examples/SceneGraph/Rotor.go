package SceneGraph

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const RotorFactoryName = "SceneGraph.Rotor"

func init() {
	NodeFactory.AddType(RotorFactoryName, reflect.TypeOf((*Rotor)(nil)).Elem())
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

func (node *Rotor) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
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

	return nil
}
