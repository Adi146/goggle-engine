package SceneGraph

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const RotorFactoryName = "SceneGraph.Rotor"

func init() {
	NodeFactory.AddType(RotorFactoryName, reflect.TypeOf((*RotorConfig)(nil)).Elem())
}

type RotorConfig struct {
	Scene.NodeConfig
	Speed float32 `yaml:"speed"`
}

func (config *RotorConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &Rotor{
		INode:  nodeBase,
		Config: config,
	}

	return node, err
}

type Rotor struct {
	Scene.INode
	Config *RotorConfig
}

func (node *Rotor) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.RotateY(node.Config.Speed * timeDelta)))

	return err
}
