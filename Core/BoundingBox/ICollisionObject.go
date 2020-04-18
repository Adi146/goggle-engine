package BoundingBox

type ICollisionObject interface {
	GetBoundingBox() AABB
	GetBoundingBoxTransformed() AABB
}
