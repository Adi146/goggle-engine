package Camera

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IFrustum interface {
	Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3)
	UpdateProjectionConfig(projectionConfig GeometryMath.PerspectiveConfig)
	Contains(volume BoundingVolume.IBoundingVolume) bool
	ContainsSphere(sphere BoundingVolume.Sphere) bool
	ContainsAABB(aabb BoundingVolume.AABB) bool
}
