package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type ChildNodeBaseConfig struct {
	NodeBaseConfig
}

func (config ChildNodeBaseConfig) Create() (INode, error) {
	return config.CreateAsChildNode()
}

func (config ChildNodeBaseConfig) CreateAsChildNode() (IChildNode, error) {
	nodeBase, err := config.NodeBaseConfig.Create()

	node := &ChildNodeBase{
		ChildNodeBaseConfig: config,
		INode:               nodeBase,
	}

	return node, err
}

type ChildNodeBase struct {
	ChildNodeBaseConfig
	INode
	parent IParentNode
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
		return node.INode.GetLocalTransformation()
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
