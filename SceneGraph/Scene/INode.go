package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type INodeConfig interface {
	Create() (INode, error)
	SetNodeID(nodeID string)
}

type INode interface {
	GetScene() *Scene
	setScene(*Scene)

	GetLocalTransformation() *Matrix.Matrix4x4
	SetLocalTransformation(*Matrix.Matrix4x4)

	GetLocalRotation() []Angle.EulerAngles
	GetLocalPosition() *Vector.Vector3

	GetNodeID() string
	GetLogFields() map[string]interface{}
}
