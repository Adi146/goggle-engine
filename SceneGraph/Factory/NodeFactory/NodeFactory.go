package NodeFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/MatrixFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

var (
	typeFactory = map[string]reflect.Type{
		"Scene.Node": reflect.TypeOf((*Scene.NodeConfig)(nil)).Elem(),
	}
)

func AddType(key string, nodeType reflect.Type) {
	typeFactory[key] = nodeType
}

type NodeConfig struct {
	Type           string                       `yaml:"type"`
	Children       map[string]NodeConfig        `yaml:"children"`
	Config         yaml.Node                    `yaml:"config"`
	Transformation []MatrixFactory.MatrixConfig `yaml:"transformation"`
}

func (config *NodeConfig) Unmarshal(nodeID string) (Scene.INode, error) {
	nodeType, ok := typeFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("node type %s is not in factory", config.Type)
	}

	nodeConfig := reflect.New(nodeType).Interface().(Scene.INodeConfig)
	if config.Config.Kind != 0 {
		config.Config.Decode(nodeConfig)
	}

	node, err := nodeConfig.Create()
	if err != nil {
		return nil, err
	}

	if err := config.UnmarshalChildren(node); err != nil {
		return nil, err
	}

	for _, transformation := range config.Transformation {
		node.SetLocalTransformation(node.GetLocalTransformation().Mul(&transformation.Matrix4x4))
	}

	return node, nil
}

func (config *NodeConfig) UnmarshalChildren(node Scene.INode) error {
	for childID, child := range config.Children {
		childNode, err := child.Unmarshal(childID)
		if err != nil {
			return err
		}

		node.AddChild(childNode)
	}

	return nil
}
