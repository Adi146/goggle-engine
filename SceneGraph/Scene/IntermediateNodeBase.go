package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type IntermediateNodeBaseConfig struct {
	ChildNodeBaseConfig
	ParentNodeBaseConfig
	NodeBaseConfig
}

func (config IntermediateNodeBaseConfig) Create() (INode, error) {
	return config.CreateAsIntermediateNode()
}

func (config IntermediateNodeBaseConfig) CreateAsIntermediateNode() (IIntermediateNode, error) {
	nodeBase, err := config.NodeBaseConfig.Create()

	node := &IntermediateNodeBase{
		ChildNodeBase: ChildNodeBase{
			ChildNodeBaseConfig: config.ChildNodeBaseConfig,
			INode:               nodeBase,
		},
		ParentNodeBase: ParentNodeBase{
			ParentNodeBaseConfig: config.ParentNodeBaseConfig,
			INode:                nodeBase,
		},
		INode: nodeBase,
	}

	return node, err
}

type IntermediateNodeBase struct {
	ChildNodeBase
	ParentNodeBase
	INode
}

func (node *IntermediateNodeBase) GetScene() *Scene {
	return node.INode.GetScene()
}

func (node *IntermediateNodeBase) setScene(scene *Scene) {
	node.ParentNodeBase.setScene(scene)
}

func (node *IntermediateNodeBase) GetLocalTransformation() *Matrix.Matrix4x4 {
	return node.INode.GetLocalTransformation()
}

func (node *IntermediateNodeBase) SetLocalTransformation(matrix *Matrix.Matrix4x4) {
	node.INode.SetLocalTransformation(matrix)
}

func (node *IntermediateNodeBase) GetLocalRotation() []Angle.EulerAngles {
	return node.INode.GetLocalRotation()
}

func (node *IntermediateNodeBase) GetLocalPosition() *Vector.Vector3 {
	return node.INode.GetLocalPosition()
}

func (node *IntermediateNodeBase) GetNodeID() string {
	return node.INode.GetNodeID()
}

func (node *IntermediateNodeBase) Tick(timeDelta float32) error {
	var err Error.ErrorCollection

	err.Push(node.ChildNodeBase.Tick(timeDelta))
	err.Push(node.ParentNodeBase.TickChildren(timeDelta))

	return err.Err()
}

func (node *IntermediateNodeBase) Draw() error {
	var err Error.ErrorCollection

	err.Push(node.ChildNodeBase.Draw())
	err.Push(node.ParentNodeBase.DrawChildren())

	return err.Err()
}
