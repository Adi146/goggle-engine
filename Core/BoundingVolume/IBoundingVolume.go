package BoundingVolume

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IBoundingVolume interface {
	GetCenter() GeometryMath.Vector3
	Transform(mat GeometryMath.Matrix4x4) IBoundingVolume
	IntersectsWith(volume IBoundingVolume) bool
}

func IntersectionAABBandAABB(box1, box2 AABB) bool {
	return (box1.Min.X() <= box2.Max.X() && box1.Max.X() >= box2.Min.X()) &&
		(box1.Min.Y() <= box2.Max.Y() && box1.Max.Y() >= box2.Min.Y()) &&
		(box1.Min.Z() <= box2.Max.Z() && box1.Max.Z() >= box2.Min.Z())
}

func IntersectionSphereAndSphere(sphere1, sphere2 Sphere) bool {
	distance := GeometryMath.Sqrt(
		(sphere1.Center.X()-sphere2.Center.X())*(sphere1.Center.X()-sphere2.Center.X()) +
			(sphere1.Center.Y()-sphere2.Center.Y())*(sphere1.Center.Y()-sphere2.Center.Y()) +
			(sphere1.Center.Z()-sphere2.Center.Z())*(sphere1.Center.Z()-sphere2.Center.Z()))
	return distance < (sphere1.Radius + sphere2.Radius)
}

func IntersectionSphereAndAABB(sphere Sphere, box AABB) bool {
	x := GeometryMath.Max(box.Min.X(), GeometryMath.Min(sphere.Center.X(), box.Max.X()))
	y := GeometryMath.Max(box.Min.Y(), GeometryMath.Min(sphere.Center.Y(), box.Max.Y()))
	z := GeometryMath.Max(box.Min.Z(), GeometryMath.Min(sphere.Center.Z(), box.Max.Z()))

	distance := GeometryMath.Sqrt(
		(x-sphere.Center.X())*(x-sphere.Center.X()) +
			(y-sphere.Center.Y())*(y-sphere.Center.Y()) +
			(z-sphere.Center.Z())*(z-sphere.Center.Z()))

	return distance < sphere.Radius
}

func IntersectionSphereAndPoint(sphere Sphere, point Point) bool {
	var distance = GeometryMath.Sqrt(
		(point.X()-sphere.Center.X())*(point.X()-sphere.Center.X()) +
			(point.Y()-sphere.Center.Y())*(point.Y()-sphere.Center.Y()) +
			(point.Z() - sphere.Center.Z()*(point.Z()) - sphere.Center.Z()))
	return distance < sphere.Radius
}

func IntersectionAABBAndPoint(box AABB, point Point) bool {
	return (point.X() >= box.Min.X() && point.X() <= box.Max.X()) &&
		(point.Y() >= box.Min.Y() && point.Y() <= box.Max.Y()) &&
		(point.Z() >= box.Min.Z() && point.Z() <= box.Max.Z())
}

func IntersectionPointAndPoint(point1, point2 Point) bool {
	return point1 == point2
}
