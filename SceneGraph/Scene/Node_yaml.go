package Scene

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
	"reflect"
)

var (
	NodeFactory = nodeFactory{
		"Scene.Node": reflect.TypeOf((*Node)(nil)).Elem(),
	}
)

type nodeFactory map[string]reflect.Type

func (factory nodeFactory) AddType(key string, nodeType reflect.Type) {
	factory[key] = nodeType
}

func (node *Node) UnmarshalYAML(value *yaml.Node) error {
	type ChildConfig struct {
		Type string `yaml:"type"`
	}

	var yamlConfig struct {
		Children       map[string]yaml.Node     `yaml:"children"`
		Transformation []GeometryMath.Matrix4x4 `yaml:"transformation"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	for _, childValue := range yamlConfig.Children {
		var childYamlConfig ChildConfig
		if err := childValue.Decode(&childYamlConfig); err != nil {
			return err
		}

		childType, ok := NodeFactory[childYamlConfig.Type]
		if !ok {
			return fmt.Errorf("node type %s is not in factory", childYamlConfig.Type)
		}

		child := reflect.New(childType).Interface().(INode)
		if err := childValue.Decode(child); err != nil {
			return err
		}
		node.AddChild(child)
	}

	if len(yamlConfig.Transformation) >= 1 || node.Transformation == nil || *node.Transformation == (GeometryMath.Matrix4x4{}) {
		node.SetLocalTransformation(GeometryMath.Identity())
		for _, transformation := range yamlConfig.Transformation {
			node.SetLocalTransformation(node.GetLocalTransformation().Mul(&transformation))
		}
	}

	return nil
}
