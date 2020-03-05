package NodeFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/MatrixFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

type Product struct {
	Scene.INode
}

type tmpProduct struct {
	Type           string                 `yaml:"type"`
	Children       map[string]Product     `yaml:"children"`
	Config         yaml.Node              `yaml:"config"`
	Transformation []MatrixFactory.Config `yaml:"transformation"`
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var tmpProduct tmpProduct
	if err := value.Decode(&tmpProduct); err != nil {
		return err
	}

	nodeType, ok := typeFactory[tmpProduct.Type]
	if !ok {
		return fmt.Errorf("node type %s is not in factory", tmpProduct.Type)
	}

	nodeConfig := reflect.New(nodeType).Interface().(Scene.INodeConfig)
	if tmpProduct.Config.Kind != 0 {
		tmpProduct.Config.Decode(nodeConfig)
	}

	node, err := nodeConfig.Create()
	if err != nil {
		return err
	}

	for _, child := range tmpProduct.Children {
		node.AddChild(child.INode)
	}

	for _, transformation := range tmpProduct.Transformation {
		node.SetLocalTransformation(node.GetLocalTransformation().Mul(&transformation.Matrix4x4))
	}

	product.INode = node
	return nil
}