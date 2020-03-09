package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type INode interface {
	AddChild(child INode)
	GetChildren() []INode

	GetParent() INode
	setParent(parent INode)

	GetScene() *Scene
	setScene(*Scene)

	GetLocalTransformation() *GeometryMath.Matrix4x4
	SetLocalTransformation(*GeometryMath.Matrix4x4)

	GetLocalRotation() []GeometryMath.EulerAngles
	GetLocalPosition() *GeometryMath.Vector3

	GetGlobalTransformation() *GeometryMath.Matrix4x4

	GetGlobalRotation() []GeometryMath.EulerAngles
	GetGlobalPosition() *GeometryMath.Vector3

	Tick(timeDelta float32) error
	Draw() error
}
