package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IChildNodeConfig interface {
	INodeConfig

	CreateAsChildNode() (IChildNode, error)
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
