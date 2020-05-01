package BoundingVolume

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type AABB struct {
	Min GeometryMath.Vector3 `yaml:"min"`
	Max GeometryMath.Vector3 `yaml:"max"`
}

func NewBoundingVolumeAABB(vertices []GeometryMath.Vector3) IBoundingVolume {
	return NewAABB(vertices)
}

func NewAABB(vertices []GeometryMath.Vector3) AABB {
	min := vertices[0]
	max := vertices[0]

	for _, vertex := range vertices[1:] {
		for i := 0; i < 3; i++ {
			max[i] = GeometryMath.Max(max[i], vertex[i])
			min[i] = GeometryMath.Min(min[i], vertex[i])
		}
	}

	return AABB{
		Min: min,
		Max: max,
	}
}

func NewDefaultAABB() AABB {
	return AABB{
		Min: GeometryMath.Vector3{-1, -1, -1},
		Max: GeometryMath.Vector3{1, 1, 1},
	}
}

func (aabb AABB) GetCenter() GeometryMath.Vector3 {
	return GeometryMath.Vector3{
		(aabb.Min[0] + aabb.Max[0]) / 2.0,
		(aabb.Min[1] + aabb.Max[1]) / 2.0,
		(aabb.Min[2] + aabb.Max[2]) / 2.0,
	}
}

func (aabb AABB) Transform(mat GeometryMath.Matrix4x4) IBoundingVolume {
	if mat == GeometryMath.Identity() {
		return aabb
	}

	vertices := []GeometryMath.Vector3{
		{aabb.Min.X(), aabb.Min.Y(), aabb.Min.Z()},
		{aabb.Min.X(), aabb.Min.Y(), aabb.Max.Z()},
		{aabb.Min.X(), aabb.Max.Y(), aabb.Max.Z()},
		{aabb.Max.X(), aabb.Max.Y(), aabb.Max.Z()},
		{aabb.Max.X(), aabb.Max.Y(), aabb.Min.Z()},
		{aabb.Max.X(), aabb.Min.Y(), aabb.Min.Z()},
		{aabb.Max.X(), aabb.Min.Y(), aabb.Max.Z()},
		{aabb.Min.X(), aabb.Max.Y(), aabb.Min.Z()},
	}

	for i := range vertices {
		vertices[i] = mat.MulVector(vertices[i])
	}

	return NewBoundingVolumeAABB(vertices)
}

func (aabb AABB) IntersectsWith(volume IBoundingVolume) bool {
	switch v := volume.(type) {
	case AABB:
		return IntersectionAABBandAABB(aabb, v)
	case Sphere:
		return IntersectionSphereAndAABB(v, aabb)
	case Point:
		return IntersectionAABBAndPoint(aabb, v)
	default:
		return false
	}
}
