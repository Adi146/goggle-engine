package Camera

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type ICamera interface {
	Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3)
	GetPosition() GeometryMath.Vector3
	GetFront() GeometryMath.Vector3
	GetUp() GeometryMath.Vector3
	GetRight() GeometryMath.Vector3
	SetProjection(projection GeometryMath.IProjectionConfig)
	GetFrustum() IFrustum
}
