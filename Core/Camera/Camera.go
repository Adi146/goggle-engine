package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type Camera struct {
	ProjectionMatrix *Matrix.Matrix4x4

	Front    *Vector.Vector3
	Up       *Vector.Vector3
	Position *Vector.Vector3

	viewMatrix *Matrix.Matrix4x4
}

func NewCamera(projectionMatrix *Matrix.Matrix4x4) *Camera {
	return &Camera{
		ProjectionMatrix: projectionMatrix,
		Front:            &Vector.Vector3{0, 0, -1},
		Up:               &Vector.Vector3{0, 1, 0},
		Position:         &Vector.Vector3{0, 0, 0},
	}
}

func NewCameraPerspective(fov float32, width float32, height float32) *Camera {
	return NewCamera(Matrix.Perspective(fov/2, width/height, 0.1, 100.0))
}

func NewCameraOrthogonal() *Camera {
	return NewCamera(Matrix.Orthogonal(-2, 2, -2, 2, -10, 100))
}

func (camera *Camera) GetViewMatrix() *Matrix.Matrix4x4 {
	return camera.viewMatrix
}

func (camera *Camera) GetProjectionMatrix() *Matrix.Matrix4x4 {
	return camera.ProjectionMatrix
}

func (camera *Camera) Tick(timeDelta float32) {
	camera.viewMatrix = Matrix.LookAt(camera.Position, camera.Position.Add(camera.Front), camera.Up)
}

func (camera *Camera) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(camera)
}

func (camera *Camera) GetPosition() *Vector.Vector3 {
	return camera.Position
}
