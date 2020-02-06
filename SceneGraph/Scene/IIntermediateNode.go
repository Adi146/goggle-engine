package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IIntermediateNodeConfig interface {
	INodeConfig

	CreateAsIntermediateNode(IIntermediateNode, error)
}

type IIntermediateNode interface {
	INode

	AddChild(child IChildNode)
	GetChildren() []IChildNode

	TickChildren(timeDelta float32) error
	DrawChildren() error

	setParent(parent IParentNode)
	GetParent() IParentNode

	GetGlobalTransformation() *Matrix.Matrix4x4

	GetGlobalRotation() []Angle.EulerAngles
	GetGlobalPosition() *Vector.Vector3

	Tick(timeDelta float32) error
	Draw() error
}
