package Camera

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type RadarFrustum struct {
	TanY          float32
	TanX          float32
	SphereFactorX float32
	SphereFactorY float32

	Position GeometryMath.Vector3
	Front    GeometryMath.Vector3
	Up       GeometryMath.Vector3
	Right    GeometryMath.Vector3

	ProjectionConfig GeometryMath.PerspectiveConfig
}

func (frustum *RadarFrustum) UpdateProjectionConfig(projectionConfig GeometryMath.PerspectiveConfig) {
	frustum.TanY = GeometryMath.Tan(GeometryMath.Radians(projectionConfig.Fovy * 0.5))
	frustum.TanX = frustum.TanY * projectionConfig.Aspect
	frustum.SphereFactorX = 1.0 / GeometryMath.Cos(GeometryMath.ATan(frustum.TanY*projectionConfig.Aspect))
	frustum.SphereFactorY = 1.0 / GeometryMath.Cos(GeometryMath.Radians(projectionConfig.Fovy*0.5))

	frustum.ProjectionConfig = projectionConfig
}

func (frustum *RadarFrustum) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	frustum.Position = position
	frustum.Front = front
	frustum.Up = up
	frustum.Right = front.Cross(up)
}

func (frustum *RadarFrustum) Contains(volume BoundingVolume.IBoundingVolume) bool {
	switch v := volume.(type) {
	case BoundingVolume.AABB:
		return frustum.ContainsAABB(v)
	case BoundingVolume.Sphere:
		return frustum.ContainsSphere(v)
	default:
		return false
	}
}

func (frustum *RadarFrustum) ContainsSphere(sphere BoundingVolume.Sphere) bool {
	v := sphere.Center.Sub(frustum.Position)

	az := v.Dot(frustum.Front)
	if az > frustum.ProjectionConfig.Far+sphere.Radius || az < frustum.ProjectionConfig.Near-sphere.Radius {
		return false
	}

	ax := v.Dot(frustum.Right)
	zz1 := az * frustum.TanX
	d1 := frustum.SphereFactorX * sphere.Radius
	if ax > zz1+d1 || ax < -zz1-d1 {
		return false
	}

	ay := v.Dot(frustum.Up)
	zz2 := az * frustum.TanY
	d2 := frustum.SphereFactorY * sphere.Radius
	if ay > zz2+d2 || ay < -zz2-d2 {
		return false
	}

	return true
}

//TODO: check correct points of box
func (frustum *RadarFrustum) ContainsAABB(aabb BoundingVolume.AABB) bool {
	pz := getVertexP(aabb, frustum.Front)
	nz := getVertexN(aabb, frustum.Front)
	pzv := pz.Sub(frustum.Position).Dot(frustum.Front)
	nzv := nz.Sub(frustum.Position).Dot(frustum.Front)
	if (pzv > frustum.ProjectionConfig.Far && nzv > frustum.ProjectionConfig.Far) || (pzv < frustum.ProjectionConfig.Near && nzv < frustum.ProjectionConfig.Near) {
		return false
	}

	dpx := getVertexP(aabb, frustum.Right).Sub(frustum.Position)
	dnx := getVertexN(aabb, frustum.Right).Sub(frustum.Position)
	pxv := dpx.Dot(frustum.Right)
	nxv := dnx.Dot(frustum.Right)
	paux := dpx.Dot(frustum.Front) * frustum.TanX
	naux := dnx.Dot(frustum.Front) * frustum.TanX
	if (pxv > paux && nxv > naux) || (pxv < -paux && nxv < -naux) {
		return false
	}

	dpy := getVertexP(aabb, frustum.Up).Sub(frustum.Position)
	dny := getVertexN(aabb, frustum.Up).Sub(frustum.Position)
	pyv := dpy.Dot(frustum.Up)
	nyv := dny.Dot(frustum.Up)
	pauy := dpy.Dot(frustum.Front) * frustum.TanY
	nauy := dny.Dot(frustum.Front) * frustum.TanY
	if (pyv > pauy && nyv > nauy) || (pyv < -pauy && nyv < -nauy) {
		return false
	}

	return true
}

func getVertexP(aabb BoundingVolume.AABB, normal GeometryMath.Vector3) GeometryMath.Vector3 {
	result := aabb.Min
	if normal.X() >= 0 {
		result[0] = aabb.Max.X()
	}
	if normal.Y() >= 0 {
		result[1] = aabb.Max.Y()
	}
	if normal.Z() >= 0 {
		result[2] = aabb.Max.Z()
	}
	return result
}

func getVertexN(aabb BoundingVolume.AABB, normal GeometryMath.Vector3) GeometryMath.Vector3 {
	result := aabb.Max
	if normal.X() >= 0 {
		result[0] = aabb.Min.X()
	}
	if normal.Y() >= 0 {
		result[1] = aabb.Min.Y()
	}
	if normal.Z() >= 0 {
		result[2] = aabb.Min.Z()
	}
	return result
}
