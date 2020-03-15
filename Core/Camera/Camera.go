package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Camera struct {
	ProjectionMatrix GeometryMath.Matrix4x4 `yaml:"projection"`
	ViewMatrix       GeometryMath.Matrix4x4 `yaml:"view"`
}

func NewCamera(projectionMatrix GeometryMath.Matrix4x4) *Camera {
	return &Camera{
		ProjectionMatrix: projectionMatrix,
		ViewMatrix:       *GeometryMath.Identity(),
	}
}

func NewCameraPerspective(fov float32, width float32, height float32) *Camera {
	return NewCamera(*GeometryMath.Perspective(fov/2, width/height, 0.1, 100.0))
}

func NewCameraOrthogonal() *Camera {
	return NewCamera(*GeometryMath.Orthogonal(-2, 2, -2, 2, -10, 100))
}

func (camera *Camera) GetViewMatrix() GeometryMath.Matrix4x4 {
	return camera.ViewMatrix
}

func (camera *Camera) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	camera.ViewMatrix = matrix
}

func (camera *Camera) GetProjectionMatrix() GeometryMath.Matrix4x4 {
	return camera.ProjectionMatrix
}

func (camera *Camera) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	camera.ProjectionMatrix = matrix
}
