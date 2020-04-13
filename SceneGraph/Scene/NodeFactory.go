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

func UnmarshalChildren(value *yaml.Node, node INode, baseTypeKey string) error {
	type ChildConfig struct {
		Type string `yaml:"type"`
	}

	var yamlConfig struct {
		Children map[string]yaml.Node `yaml:"children"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	for _, childValue := range yamlConfig.Children {
		var childYamlConfig ChildConfig
		if err := childValue.Decode(&childYamlConfig); err != nil {
			return err
		}

		child, err := NodeFactory.Get(childYamlConfig.Type)
		if err != nil {
			return err
		}

		if asSubNode, isSubNode := child.(ISubNode); isSubNode {
			base, err := NodeFactory.Get(baseTypeKey)
			if err != nil {
				return err
			}

			asSubNode.SetBase(base)
		}

		if err := childValue.Decode(child); err != nil {
			return err
		}

		AddChild(node, child)
	}

	return nil
}

func UnmarshalBase(value *yaml.Node, node ISubNode) error {
	var yamlConfig struct {
		BaseValue yaml.Node `yaml:"baseType"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	base := node.GetBase()
	if yamlConfig.BaseValue.Kind == yaml.ScalarNode {
		var baseKey string
		if err := yamlConfig.BaseValue.Decode(&baseKey); err != nil {
			return err
		}

		tmpBase, err := NodeFactory.Get(baseKey)
		if err != nil {
			return err
		}
		base = tmpBase
	} else if base == nil {
		base = &Node{}
	}
	if err := value.Decode(base); err != nil {
		return err
	}
	node.SetBase(base)

	return nil
}

func AddChild(parent INode, child INode) {
	parent.AddChild(child)
	child.SetParent(parent)
}
