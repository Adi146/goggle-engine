package YamlFactory

import (
	"fmt"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
)

var NodeFactory = map[string]reflect.Type{
	"Scene.ChildBaseNode":        reflect.TypeOf((*Scene.ChildNodeBaseConfig)(nil)).Elem(),
	"Scene.ParentBaseNode":       reflect.TypeOf((*Scene.ParentNodeBaseConfig)(nil)).Elem(),
	"Scene.IntermediateNodeBase": reflect.TypeOf((*Scene.IntermediateNodeBaseConfig)(nil)).Elem(),
}

type NodeConfig struct {
	Type           string                 `yaml:"type"`
	Children       map[string]NodeConfig  `yaml:"children"`
	Config         yaml.Node              `yaml:"config"`
	Transformation []map[string]yaml.Node `yaml:"transformation"`
}

func (config *NodeConfig) Unmarshal(nodeID string) (Scene.INode, error) {
	nodeType, ok := NodeFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("node type %s is not in factory", config.Type)
	}

	nodeConfig := reflect.New(nodeType).Interface().(Scene.INodeConfig)
	if config.Config.Kind != 0 {
		nodeConfig.SetNodeID(nodeID)
		config.Config.Decode(nodeConfig)
	}

	node, err := nodeConfig.Create()
	if err != nil {
		return nil, err
	}

	if err := config.UnmarshalChildren(node); err != nil {
		return nil, err
	}

	if err := config.UnmarshalTransformation(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (config *NodeConfig) UnmarshalChildren(node Scene.INode) error {
	if nodeAsParent, isParent := node.(Scene.IParentNode); isParent {
		for childID, child := range config.Children {
			childNode, err := child.Unmarshal(childID)
			if err != nil {
				return err
			}

			if childNodeAsChild, isChild := childNode.(Scene.IChildNode); isChild {
				nodeAsParent.AddChild(childNodeAsChild)
			} else {
				return fmt.Errorf("node type %s is not a IChildNode", config.Type)
			}
		}
	}

	return nil
}

func (config *NodeConfig) UnmarshalTransformation(node Scene.INode) error {
	for _, transformationGroup := range config.Transformation {
		for transformationType, transformationConfig := range transformationGroup {
			matrixType, ok := MatrixFactory[transformationType]
			if !ok {
				return fmt.Errorf("matrix type %s not in factory", transformationType)
			}

			matrixConfig := reflect.New(matrixType).Interface().(IYamlMatrixConfig)
			transformationConfig.Decode(matrixConfig)

			node.SetLocalTransformation(node.GetLocalTransformation().Mul(matrixConfig.Decode()))
		}
	}

	return nil
}
