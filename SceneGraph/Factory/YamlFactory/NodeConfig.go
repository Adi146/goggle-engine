package YamlFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

var NodeFactory = map[string]reflect.Type{
	"Scene.ChildBaseNode":        reflect.TypeOf((*Scene.ChildNodeBase)(nil)).Elem(),
	"Scene.ParentBaseNode":       reflect.TypeOf((*Scene.ParentNodeBase)(nil)).Elem(),
	"Scene.IntermediateNodeBase": reflect.TypeOf((*Scene.IntermediateNodeBase)(nil)).Elem(),
}

type NodeConfig struct {
	Type           string                 `yaml:"type"`
	Children       map[string]NodeConfig  `yaml:"children"`
	Config         yaml.Node              `yaml:"config"`
	Transformation []map[string]yaml.Node `yaml:"transformation"`
}

func (nodeConfig *NodeConfig) Unmarshal(nodeID string) (Scene.INode, error) {
	nodeType, ok := NodeFactory[nodeConfig.Type]
	if !ok {
		return nil, fmt.Errorf("node type %s is not in factory", nodeConfig.Type)
	}

	node := reflect.New(nodeType).Interface().(Scene.INode)
	if nodeConfig.Config.Kind != 0 {
		nodeConfig.Config.Decode(node)
	}

	if err := node.Init(nodeID); err != nil {
		return nil, err
	}

	if err := nodeConfig.UnmarshalChildren(node); err != nil {
		return nil, err
	}

	if err := nodeConfig.UnmarshalTransformation(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (nodeConfig *NodeConfig) UnmarshalChildren(node Scene.INode) error {
	if nodeAsParent, isParent := node.(Scene.IParentNode); isParent {
		for childID, child := range nodeConfig.Children {
			childNode, err := child.Unmarshal(childID)
			if err != nil {
				return err
			}

			if childNodeAsChild, isChild := childNode.(Scene.IChildNode); isChild {
				nodeAsParent.AddChild(childNodeAsChild)
			} else {
				return fmt.Errorf("node type %s is not a IChildNode", nodeConfig.Type)
			}
		}
	}

	return nil
}

func (nodeConfig *NodeConfig) UnmarshalTransformation(node Scene.INode) error {
	for _, transformationGroup := range nodeConfig.Transformation {
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
