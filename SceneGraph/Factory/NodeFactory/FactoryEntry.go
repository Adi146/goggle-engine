package NodeFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/MatrixFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

type FactoryEntry struct {
	Type           string                  `yaml:"type"`
	Children       map[string]FactoryEntry `yaml:"children"`
	Config         yaml.Node               `yaml:"config"`
	Transformation []MatrixFactory.Config  `yaml:"transformation"`
}

func (entry *FactoryEntry) Unmarshal(nodeID string) (Scene.INode, error) {
	nodeType, ok := typeFactory[entry.Type]
	if !ok {
		return nil, fmt.Errorf("node type %s is not in factory", entry.Type)
	}

	nodeConfig := reflect.New(nodeType).Interface().(Scene.INodeConfig)
	if entry.Config.Kind != 0 {
		entry.Config.Decode(nodeConfig)
	}

	node, err := nodeConfig.Create()
	if err != nil {
		return nil, err
	}

	if err := entry.UnmarshalChildren(node); err != nil {
		return nil, err
	}

	for _, transformation := range entry.Transformation {
		node.SetLocalTransformation(node.GetLocalTransformation().Mul(&transformation.Matrix4x4))
	}

	return node, nil
}

func (entry *FactoryEntry) UnmarshalChildren(node Scene.INode) error {
	for childID, child := range entry.Children {
		childNode, err := child.Unmarshal(childID)
		if err != nil {
			return err
		}

		node.AddChild(childNode)
	}

	return nil
}
