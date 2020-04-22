package Camera

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type PlaneFrustum struct {
	Dimensions struct {
		FarHeight  float32
		FarWidth   float32
		NearHeight float32
		NearWidth  float32
	}

	Planes struct {
		Far    GeometryMath.Plane3
		Near   GeometryMath.Plane3
		Top    GeometryMath.Plane3
		Bottom GeometryMath.Plane3
		Right  GeometryMath.Plane3
		Left   GeometryMath.Plane3
	}

	ProjectionConfig GeometryMath.PerspectiveConfig
}

func (frustum *PlaneFrustum) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	right := front.Cross(up)

	centerFar := position.Add(front.MulScalar(frustum.ProjectionConfig.Far))
	centerNear := position.Add(front.MulScalar(frustum.ProjectionConfig.Near))

	frustum.Planes.Near = GeometryMath.NewPlane(front, centerNear)
	frustum.Planes.Far = GeometryMath.NewPlane(front.Invert(), centerFar)

	nearTop := centerNear.Add(up.MulScalar(frustum.Dimensions.NearHeight))
	nearBottom := centerNear.Sub(up.MulScalar(frustum.Dimensions.NearHeight))
	nearRight := centerNear.Add(right.MulScalar(frustum.Dimensions.NearWidth))
	nearLeft := centerNear.Sub(right.MulScalar(frustum.Dimensions.NearWidth))

	aux := nearTop.Sub(position).Normalize()
	normal := aux.Cross(right)
	frustum.Planes.Top = GeometryMath.NewPlane(normal, nearTop)

	aux = nearBottom.Sub(position).Normalize()
	normal = right.Cross(aux)
	frustum.Planes.Bottom = GeometryMath.NewPlane(normal, nearBottom)

	aux = nearRight.Sub(position).Normalize()
	normal = up.Cross(aux)
	frustum.Planes.Right = GeometryMath.NewPlane(normal, nearRight)

	aux = nearLeft.Sub(position).Normalize()
	normal = aux.Cross(up)
	frustum.Planes.Left = GeometryMath.NewPlane(normal, nearLeft)
}

func (frustum *PlaneFrustum) UpdateProjectionConfig(projectionConfig GeometryMath.PerspectiveConfig) {
	frustum.ProjectionConfig = projectionConfig

	tanY := GeometryMath.Tan(GeometryMath.Radians(projectionConfig.Fovy * 0.5))

	frustum.Dimensions.FarHeight = frustum.ProjectionConfig.Far * tanY
	frustum.Dimensions.NearHeight = frustum.ProjectionConfig.Near * tanY
	frustum.Dimensions.FarWidth = frustum.Dimensions.FarHeight * projectionConfig.Aspect
	frustum.Dimensions.NearWidth = frustum.Dimensions.NearHeight * projectionConfig.Aspect
}

func (frustum *PlaneFrustum) Contains(volume BoundingVolume.IBoundingVolume) bool {
	switch v := volume.(type) {
	case BoundingVolume.AABB:
		return frustum.ContainsAABB(v)
	case BoundingVolume.Sphere:
		return frustum.ContainsSphere(v)
	default:
		return false
	}
}

func (frustum *PlaneFrustum) ContainsSphere(sphere BoundingVolume.Sphere) bool {
	for _, plane := range frustum.GetAllPlanes() {
		if plane.Distance(sphere.Center) < -sphere.Radius {
			return false
		}
	}

	return true
}

func (frustum *PlaneFrustum) ContainsAABB(aabb BoundingVolume.AABB) bool {
	for _, plane := range frustum.GetAllPlanes() {
		if plane.Distance(getVertexP(aabb, plane.Normal)) < 0 {
			return false
		}
	}

	return true
}

func (frustum *PlaneFrustum) GetAllPlanes() []GeometryMath.Plane3 {
	return []GeometryMath.Plane3{
		frustum.Planes.Far,
		frustum.Planes.Near,
		frustum.Planes.Right,
		frustum.Planes.Left,
		frustum.Planes.Top,
		frustum.Planes.Bottom,
	}
}
