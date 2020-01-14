package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IntermediateNodeBase struct {
	ChildNodeBase
	ParentNodeBase
	*NodeBase
}

func NewIntermediateNodeBase() *IntermediateNodeBase {
	baseNode := NewNodeBase()

	return &IntermediateNodeBase{
		ChildNodeBase: ChildNodeBase{
			NodeBase: baseNode,
			parent:   nil,
		},
		ParentNodeBase: ParentNodeBase{
			NodeBase: baseNode,
			children: []IChildNode{},
		},
		NodeBase: baseNode,
	}
}

func (node *IntermediateNodeBase) Init() error {

	if node.NodeBase == nil {
		baseNode := NewNodeBase()

		node.ChildNodeBase.NodeBase = baseNode
		node.ParentNodeBase.NodeBase = baseNode
		node.NodeBase = baseNode
	}

	return nil
}

func (node *IntermediateNodeBase) GetScene() *Scene {
	return node.NodeBase.GetScene()
}

func (node *IntermediateNodeBase) setScene(scene *Scene) {
	node.ParentNodeBase.setScene(scene)
}

func (node *IntermediateNodeBase) GetLocalTransformation() *Matrix.Matrix4x4 {
	return node.NodeBase.GetLocalTransformation()
}

func (node *IntermediateNodeBase) SetLocalTransformation(matrix *Matrix.Matrix4x4) {
	node.NodeBase.SetLocalTransformation(matrix)
}

func (node *IntermediateNodeBase) GetLocalRotation() []Angle.EulerAngles {
	return node.NodeBase.GetLocalRotation()
}

func (node *IntermediateNodeBase) GetLocalPosition() *Vector.Vector3 {
	return node.NodeBase.GetLocalPosition()
}
