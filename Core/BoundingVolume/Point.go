package BoundingVolume

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type Point GeometryMath.Vector3

func NewDefaultPoint() Point {
	return Point{0, 0, 0}
}

func (point Point) X() float32 {
	return GeometryMath.Vector3(point).X()
}

func (point Point) Y() float32 {
	return GeometryMath.Vector3(point).Y()
}

func (point Point) Z() float32 {
	return GeometryMath.Vector3(point).Z()
}

func (point Point) GetCenter() GeometryMath.Vector3 {
	return GeometryMath.Vector3(point)
}

func (point Point) Transform(mat GeometryMath.Matrix4x4) IBoundingVolume {
	if mat == GeometryMath.Identity() {
		return point
	}

	return Point(mat.MulVector(GeometryMath.Vector3(point)))
}

func (point Point) IntersectsWith(volume IBoundingVolume) bool {
	switch v := volume.(type) {
	case Point:
		return IntersectionPointAndPoint(v, point)
	case AABB:
		return IntersectionAABBAndPoint(v, point)
	case Sphere:
		return IntersectionSphereAndPoint(v, point)
	default:
		return false
	}
}
