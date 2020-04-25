package Camera

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type PlaneFrustum struct {
	Planes struct {
		Far    GeometryMath.Plane3
		Near   GeometryMath.Plane3
		Top    GeometryMath.Plane3
		Bottom GeometryMath.Plane3
		Right  GeometryMath.Plane3
		Left   GeometryMath.Plane3
	}

	Dimensions struct {
		FarHeight  float32
		FarWidth   float32
		NearHeight float32
		NearWidth  float32
	}

	ProjectionConfig GeometryMath.IProjectionConfig
}

func (frustum *PlaneFrustum) Update(camera ICamera) {
	position := camera.GetPosition()
	front := camera.GetFront()
	up := camera.GetUp()
	right := camera.GetRight()

	centerFar := position.Add(front.MulScalar(frustum.ProjectionConfig.GetFar()))
	centerNear := position.Add(front.MulScalar(frustum.ProjectionConfig.GetNear()))

	frustum.Planes.Near = GeometryMath.NewPlane(front, centerNear)
	frustum.Planes.Far = GeometryMath.NewPlane(front.Invert(), centerFar)

	nearTop := centerNear.Add(up.MulScalar(frustum.Dimensions.NearHeight))
	nearBottom := centerNear.Sub(up.MulScalar(frustum.Dimensions.NearHeight))
	nearRight := centerNear.Add(right.MulScalar(frustum.Dimensions.NearWidth))
	nearLeft := centerNear.Sub(right.MulScalar(frustum.Dimensions.NearWidth))

	farTop := centerFar.Add(up.MulScalar(frustum.Dimensions.FarHeight))
	farBottom := centerFar.Sub(up.MulScalar(frustum.Dimensions.FarHeight))
	farRight := centerFar.Add(right.MulScalar(frustum.Dimensions.FarWidth))
	farLeft := centerFar.Sub(right.MulScalar(frustum.Dimensions.FarWidth))

	aux := farTop.Sub(nearTop).Normalize()
	normal := aux.Cross(right)
	frustum.Planes.Top = GeometryMath.NewPlane(normal, nearTop)

	aux = farBottom.Sub(nearBottom).Normalize()
	normal = right.Cross(aux)
	frustum.Planes.Bottom = GeometryMath.NewPlane(normal, nearBottom)

	aux = farRight.Sub(nearRight).Normalize()
	normal = up.Cross(aux)
	frustum.Planes.Right = GeometryMath.NewPlane(normal, nearRight)

	aux = farLeft.Sub(nearLeft).Normalize()
	normal = aux.Cross(up)
	frustum.Planes.Left = GeometryMath.NewPlane(normal, nearLeft)
}

func (frustum *PlaneFrustum) UpdateProjectionConfig(projectionConfig GeometryMath.IProjectionConfig) {
	frustum.ProjectionConfig = projectionConfig

	frustum.Dimensions.NearWidth, frustum.Dimensions.NearHeight = projectionConfig.GetPlane(projectionConfig.GetNear())
	frustum.Dimensions.FarWidth, frustum.Dimensions.FarHeight = projectionConfig.GetPlane(projectionConfig.GetFar())
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

func (frustum *PlaneFrustum) GetProjectionConfig() GeometryMath.IProjectionConfig {
	return frustum.ProjectionConfig
}
