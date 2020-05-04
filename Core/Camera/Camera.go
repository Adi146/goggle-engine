package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Camera struct {
	position GeometryMath.Vector3

	front GeometryMath.Vector3
	up    GeometryMath.Vector3
	right GeometryMath.Vector3

	projectionMatrix GeometryMath.Matrix4x4
	viewMatrix       GeometryMath.Matrix4x4

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

func (camera *Camera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.position = position

	camera.front = front
	camera.up = up
	camera.right = front.Cross(up)

	camera.viewMatrix = GeometryMath.LookAt(position, position.Add(front), up)

	camera.frustum.Update(camera)
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.projectionMatrix = projection.Decode()

	camera.frustum.UpdateProjectionConfig(projection)
}

func (camera *Camera) GetProjectionMatrix() GeometryMath.Matrix4x4 {
	return camera.projectionMatrix
}

func (camera *Camera) GetViewMatrix() GeometryMath.Matrix4x4 {
	return camera.viewMatrix
}

func (camera *Camera) GetFrustum() IFrustum {
	return &camera.frustum
}
