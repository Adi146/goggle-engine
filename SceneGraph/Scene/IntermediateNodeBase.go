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

func (node *IntermediateNodeBase) Init(nodeID string) error {
	if node.NodeBase == nil {
		node.NodeBase = &NodeBase{
			scene:          nil,
			transformation: Matrix.Identity(),
		}
		if err := node.NodeBase.Init(nodeID); err != nil {
			return err
		}

		node.ChildNodeBase.NodeBase = node.NodeBase
		node.ParentNodeBase.NodeBase = node.NodeBase
		if err := node.ChildNodeBase.Init(nodeID); err != nil {
			return err
		}
		if err := node.ParentNodeBase.Init(nodeID); err != nil {
			return err
		}
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

func (node *IntermediateNodeBase) GetNodeID() string {
	return node.NodeBase.GetNodeID()
}
