package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type ChildNodeBase struct {
	*NodeBase
	parent IParentNode
}

func (node *ChildNodeBase) Init(nodeID string) error {
	if node.NodeBase == nil {
		node.NodeBase = &NodeBase{
			scene:          nil,
			transformation: Matrix.Identity(),
		}
		if err := node.NodeBase.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *ChildNodeBase) GetParent() IParentNode {
	return node.parent
}

func (node *ChildNodeBase) setParent(parent IParentNode) {
	node.parent = parent
	if parent != nil {
		node.setScene(parent.GetScene())
	} else {
		node.setScene(nil)
	}
}

func (node *ChildNodeBase) GetGlobalTransformation() *Matrix.Matrix4x4 {
	if node.GetParent() == nil {
		return node.transformation
	} else {
		if parentAsChild, isChild := node.GetParent().(IChildNode); isChild {
			return parentAsChild.GetGlobalTransformation().Mul(node.GetLocalTransformation())
		} else {
			return node.GetParent().GetLocalTransformation().Mul(node.GetLocalTransformation())
		}
	}
}

func (node *ChildNodeBase) GetGlobalRotation() []Angle.EulerAngles {
	return Angle.ExtractFromMatrix(node.GetGlobalTransformation())
}

func (node *ChildNodeBase) GetGlobalPosition() *Vector.Vector3 {
	return node.GetGlobalTransformation().MulVector(&Vector.Vector3{0, 0, 0})
}

func (node *ChildNodeBase) Tick(timeDelta float32) error {
	return nil
}

func (node *ChildNodeBase) Draw() error {
	return nil
}
