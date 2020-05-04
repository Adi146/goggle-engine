package PointLight

import (
	core "github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Camera struct {
	core.Camera
	ViewProjectionMatrices [6]GeometryMath.Matrix4x4
	ProjectionMatrix       GeometryMath.Matrix4x4
}

func (camera *Camera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.ViewProjectionMatrices[0] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0}))
	camera.ViewProjectionMatrices[1] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{-1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0}))
	camera.ViewProjectionMatrices[2] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, 1.0}))
	camera.ViewProjectionMatrices[3] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, -1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, -1.0}))
	camera.ViewProjectionMatrices[4] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 0.0, 1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0}))
	camera.ViewProjectionMatrices[5] = camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 0.0, -1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0}))

	camera.Camera.Update(position, front, up)
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.ProjectionMatrix = projection.Decode()

	frustumProjection := GeometryMath.OrthographicConfig{
		Left:   -projection.GetFar(),
		Right:  projection.GetFar(),
		Bottom: -projection.GetFar(),
		Top:    projection.GetFar(),
		Near:   -projection.GetFar(),
		Far:    projection.GetFar(),
	}
	camera.Camera.SetProjection(&frustumProjection)
}
