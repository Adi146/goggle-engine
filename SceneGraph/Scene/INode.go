package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type INode interface {
	Init(nodeID string) error

	GetScene() *Scene
	setScene(*Scene)

	GetLocalTransformation() *Matrix.Matrix4x4
	SetLocalTransformation(*Matrix.Matrix4x4)

	GetLocalRotation() []Angle.EulerAngles
	GetLocalPosition() *Vector.Vector3

	GetNodeID() string
	GetLogFields() map[string]interface{}
}

type IParentNode interface {
	INode

	AddChild(child IChildNode)
	GetChildren() []IChildNode

	TickChildren(timeDelta float32) error
	DrawChildren() error
}

type IChildNode interface {
	INode

	setParent(parent IParentNode)
	GetParent() IParentNode

	GetGlobalTransformation() *Matrix.Matrix4x4

	GetGlobalRotation() []Angle.EulerAngles
	GetGlobalPosition() *Vector.Vector3

	Tick(timeDelta float32) error
	Draw() error
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
