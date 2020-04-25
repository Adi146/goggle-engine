package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Camera struct {
	position GeometryMath.Vector3

	front GeometryMath.Vector3
	up    GeometryMath.Vector3
	right GeometryMath.Vector3

	frustum PlaneFrustum
}

func (camera *Camera) GetPosition() GeometryMath.Vector3 {
	return camera.position
}

func (camera *Camera) GetFront() GeometryMath.Vector3 {
	return camera.front
}

func (camera *Camera) GetUp() GeometryMath.Vector3 {
	return camera.up
}

func (camera *Camera) GetRight() GeometryMath.Vector3 {
	return camera.right
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.frustum.UpdateProjectionConfig(projection)
}

func (camera *Camera) GetFrustum() IFrustum {
	return &camera.frustum
}
