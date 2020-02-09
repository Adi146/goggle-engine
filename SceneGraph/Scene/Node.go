package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type Node struct {
	scene          *Scene
	transformation *Matrix.Matrix4x4

	children []INode
	parent   INode

	Config *NodeConfig
}

func (node *Node) AddChild(child INode) {
	node.children = append(node.children, child)
	child.setParent(node)
}

func (node *Node) GetChildren() []INode {
	return node.children
}

func (node *Node) GetParent() INode {
	return node.parent
}

func (node *Node) setParent(parent INode) {
	node.parent = parent
	if parent != nil {
		node.setScene(parent.GetScene())
	} else {
		node.setScene(nil)
	}
}

func (node *Node) GetScene() *Scene {
	return node.scene
}

func (node *Node) setScene(scene *Scene) {
	node.scene = scene

	for _, child := range node.children {
		child.setScene(scene)
	}
}

func (node *Node) GetLocalTransformation() *Matrix.Matrix4x4 {
	return node.transformation
}

func (node *Node) SetLocalTransformation(matrix *Matrix.Matrix4x4) {
	node.transformation = matrix
}

func (node *Node) GetLocalRotation() []Angle.EulerAngles {
	return Angle.ExtractFromMatrix(node.GetLocalTransformation())
}

func (node *Node) GetLocalPosition() *Vector.Vector3 {
	return node.GetLocalTransformation().MulVector(&Vector.Vector3{0, 0, 0})
}

func (node *Node) GetGlobalTransformation() *Matrix.Matrix4x4 {
	if parent := node.GetParent(); parent == nil {
		return node.GetLocalTransformation()
	} else {
		return node.parent.GetGlobalTransformation().Mul(node.GetLocalTransformation())
	}
}

func (node *Node) GetGlobalRotation() []Angle.EulerAngles {
	return Angle.ExtractFromMatrix(node.GetGlobalTransformation())
}

func (node *Node) GetGlobalPosition() *Vector.Vector3 {
	return node.GetGlobalTransformation().MulVector(&Vector.Vector3{0, 0, 0})
}

func (node *Node) Tick(timeDelta float32) error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Tick(timeDelta))
	}

	return err.Err()
}

func (node *Node) Draw() error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Draw())
	}

	return err.Err()
}
