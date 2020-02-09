package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type INode interface {
	AddChild(child INode)
	GetChildren() []INode

	GetParent() INode
	setParent(parent INode)

	GetScene() *Scene
	setScene(*Scene)

	GetLocalTransformation() *Matrix.Matrix4x4
	SetLocalTransformation(*Matrix.Matrix4x4)

	GetLocalRotation() []Angle.EulerAngles
	GetLocalPosition() *Vector.Vector3

	GetGlobalTransformation() *Matrix.Matrix4x4

	GetGlobalRotation() []Angle.EulerAngles
	GetGlobalPosition() *Vector.Vector3

	Tick(timeDelta float32) error
	Draw() error
}
