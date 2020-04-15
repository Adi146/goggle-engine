package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type INode interface {
	AddChild(child INode, id string)
	GetChildren() map[string]INode
	GetGrandChildById(id string) INode

	GetParent() INode
	SetParent(parent INode)

	GetScene() *Scene
	SetScene(*Scene)

	GetLocalTransformation() *GeometryMath.Matrix4x4
	SetLocalTransformation(*GeometryMath.Matrix4x4)

	GetLocalRotation() []GeometryMath.EulerAngles
	GetLocalPosition() *GeometryMath.Vector3

	GetGlobalTransformation() *GeometryMath.Matrix4x4

	GetGlobalRotation() []GeometryMath.EulerAngles
	GetGlobalPosition() *GeometryMath.Vector3

	Tick(timeDelta float32) error
	GetBase() INode
}
