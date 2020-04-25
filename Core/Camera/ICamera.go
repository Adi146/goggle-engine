package Camera

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type ICamera interface {
	GetPosition() GeometryMath.Vector3
	GetFront() GeometryMath.Vector3
	GetUp() GeometryMath.Vector3
	GetRight() GeometryMath.Vector3
	SetProjection(projection GeometryMath.IProjectionConfig)
	GetFrustum() IFrustum
}
