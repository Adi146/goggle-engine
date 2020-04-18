package BoundingBox

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type AABB struct {
	Min GeometryMath.Vector3
	Max GeometryMath.Vector3
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

func (aabb *AABB) GetCenter() GeometryMath.Vector3 {
	return GeometryMath.Vector3{
		(aabb.Min[0] + aabb.Max[0]) / 2.0,
		(aabb.Min[1] + aabb.Max[1]) / 2.0,
		(aabb.Min[2] + aabb.Max[2]) / 2.0,
	}
}

func (aabb *AABB) Transform(mat GeometryMath.Matrix4x4) AABB {
	if mat == GeometryMath.Identity() {
		return *aabb
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

	return NewAABB(vertices)
}

func (aabb *AABB) IntersectsWith(other AABB) bool {
	return (aabb.Min.X() <= other.Max.X() && aabb.Max.X() >= other.Min.X()) &&
		(aabb.Min.Y() <= other.Max.Y() && aabb.Max.Y() >= other.Min.Y()) &&
		(aabb.Min.Z() <= other.Max.Z() && aabb.Max.Z() >= other.Min.Z())
}
