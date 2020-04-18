package Scene

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
)

var (
	NodeFactory = nodeFactory{}
)

type nodeFactory map[string]reflect.Type

func (factory nodeFactory) AddType(key string, nodeType reflect.Type) {
	factory[key] = nodeType
}

func (factory nodeFactory) Get(key string) (INode, error) {
	nodeType, ok := NodeFactory[key]
	if !ok {
		return nil, fmt.Errorf("node type %s is not in factory", key)
	}

	return reflect.New(nodeType).Interface().(INode), nil
}

func UnmarshalChildren(value *yaml.Node, node INode) error {
	type ChildConfig struct {
		Type string `yaml:"type"`
	}

	var yamlConfig struct {
		Children map[string]yaml.Node `yaml:"children"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	for id, childValue := range yamlConfig.Children {
		var childYamlConfig ChildConfig
		if err := childValue.Decode(&childYamlConfig); err != nil {
			return err
		}

		child, err := NodeFactory.Get(childYamlConfig.Type)
		if err != nil {
			return err
		}

		if asSubNode, isSubNode := child.(ISubNode); isSubNode {
			base := Node{}
			if err := childValue.Decode(&base); err != nil {
				return err
			}

			asSubNode.SetBase(&base)
		}

		AddChild(node, child, id)
		if err := childValue.Decode(child); err != nil {
			return err
		}
	}

	return nil
}

func AddChild(parent INode, child INode, childID string) {
	parent.AddChild(child, childID)
	child.SetParent(parent, childID)
}
