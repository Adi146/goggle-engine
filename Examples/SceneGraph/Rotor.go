package SceneGraph

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type Rotor struct {
	Scene.IIntermediateNode
	Speed float32 `yaml:"speed"`
}

func init() {
	Factory.NodeFactory["SceneGraph.Rotor"] = reflect.TypeOf((*Rotor)(nil)).Elem()
}

func (node *Rotor) Init(nodeID string) error {
	if node.IIntermediateNode == nil {
		node.IIntermediateNode = &Scene.IntermediateNodeBase{}
		if err := node.IIntermediateNode.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *Rotor) Tick(timeDelta float32) error {
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.RotateY(node.Speed * timeDelta)))

	return nil
}
