package BoundingVolume

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Sphere struct {
	Center GeometryMath.Vector3
	Radius float32
}

func NewBoundingVolumeSphere(vertices []GeometryMath.Vector3) IBoundingVolume {
	return NewSphere(vertices)
}

func NewSphere(vertices []GeometryMath.Vector3) Sphere {
	aabb := NewBoundingVolumeAABB(vertices)
	center := aabb.GetCenter()

	return NewSphereWithCenter(center, vertices)
}

func NewSphereWithCenter(center GeometryMath.Vector3, vertices []GeometryMath.Vector3) Sphere {
	radius := float32(0.0)
	for _, vertex := range vertices {
		radius = GeometryMath.Max(radius, vertex.Sub(center).Length())
	}

	return Sphere{
		Center: center,
		Radius: radius,
	}
}

func (sphere Sphere) GetCenter() GeometryMath.Vector3 {
	return sphere.Center
}

func (sphere Sphere) Transform(mat GeometryMath.Matrix4x4) IBoundingVolume {
	vertices := []GeometryMath.Vector3{
		mat.MulVector(sphere.Center.Add(GeometryMath.Vector3{1, 0, 0}.MulScalar(sphere.Radius))),
		mat.MulVector(sphere.Center.Add(GeometryMath.Vector3{0, 1, 0}.MulScalar(sphere.Radius))),
		mat.MulVector(sphere.Center.Add(GeometryMath.Vector3{0, 0, 1}.MulScalar(sphere.Radius))),
	}

	return NewSphereWithCenter(mat.MulVector(sphere.Center), vertices)
}

func (sphere Sphere) IntersectsWith(volume IBoundingVolume) bool {
	switch v := volume.(type) {
	case AABB:
		return IntersectionSphereAndAABB(sphere, v)
	case Sphere:
		return IntersectionSphereAndSphere(sphere, v)
	default:
		return false
	}
}
