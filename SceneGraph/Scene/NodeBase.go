package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type NodeBase struct {
	scene          *Scene
	transformation *Matrix.Matrix4x4
	nodeID         string
}

func (node *NodeBase) Init(nodeID string) error {
	node.nodeID = nodeID
	return nil
}

func (node *NodeBase) GetScene() *Scene {
	return node.scene
}

func (node *NodeBase) setScene(scene *Scene) {
	node.scene = scene
}

func (node *NodeBase) GetLocalTransformation() *Matrix.Matrix4x4 {
	return node.transformation
}

func (node *NodeBase) SetLocalTransformation(matrix *Matrix.Matrix4x4) {
	node.transformation = matrix
}

func (node *NodeBase) GetLocalRotation() []Angle.EulerAngles {
	return Angle.ExtractFromMatrix(node.GetLocalTransformation())
}

func (node *NodeBase) GetLocalPosition() *Vector.Vector3 {
	return node.GetLocalTransformation().MulVector(&Vector.Vector3{0, 0, 0})
}

func (node *NodeBase) GetNodeID() string {
	return node.nodeID
}
