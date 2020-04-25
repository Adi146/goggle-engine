package Camera

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IFrustum interface {
	Update(camera ICamera)
	UpdateProjectionConfig(projectionConfig GeometryMath.IProjectionConfig)
	Contains(volume BoundingVolume.IBoundingVolume) bool
	ContainsSphere(sphere BoundingVolume.Sphere) bool
	ContainsAABB(aabb BoundingVolume.AABB) bool
	GetProjectionConfig() GeometryMath.IProjectionConfig
}
