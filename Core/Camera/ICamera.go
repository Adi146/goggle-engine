package Camera

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type ICamera interface {
	GetViewMatrix() GeometryMath.Matrix4x4
	GetPosition() GeometryMath.Vector3
	GetFront() GeometryMath.Vector3
	GetUp() GeometryMath.Vector3
	SetProjection(projection GeometryMath.PerspectiveConfig)
	GetProjection() GeometryMath.PerspectiveConfig
	GetProjectionMatrix() GeometryMath.Matrix4x4
}
