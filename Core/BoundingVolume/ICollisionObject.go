package BoundingVolume

type ICollisionObject interface {
	GetBoundingVolume() IBoundingVolume
	GetBoundingVolumeTransformed() IBoundingVolume
}
