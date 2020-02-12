package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
)

type Camera struct {
	ProjectionMatrix Matrix.Matrix4x4
	ViewMatrix       Matrix.Matrix4x4
}

func NewCamera(projectionMatrix Matrix.Matrix4x4) *Camera {
	return &Camera{
		ProjectionMatrix: projectionMatrix,
		ViewMatrix:       *Matrix.Identity(),
	}
}

func NewCameraPerspective(fov float32, width float32, height float32) *Camera {
	return NewCamera(*Matrix.Perspective(fov/2, width/height, 0.1, 100.0))
}

func NewCameraOrthogonal() *Camera {
	return NewCamera(*Matrix.Orthogonal(-2, 2, -2, 2, -10, 100))
}

func (camera *Camera) Get() Camera {
	return *camera
}

func (camera *Camera) Set(val Camera) {
	camera = &val
}

func (camera *Camera) GetViewMatrix() Matrix.Matrix4x4 {
	return camera.ViewMatrix
}

func (camera *Camera) SetViewMatrix(matrix Matrix.Matrix4x4) {
	camera.ViewMatrix = matrix
}

func (camera *Camera) GetProjectionMatrix() Matrix.Matrix4x4 {
	return camera.ProjectionMatrix
}

func (camera *Camera) SetProjectionMatrix(matrix Matrix.Matrix4x4) {
	camera.ProjectionMatrix = matrix
}
