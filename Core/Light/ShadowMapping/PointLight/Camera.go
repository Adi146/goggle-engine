package PointLight

import (
	core "github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
)

const (
	camera_size_section = 384
)

type Camera struct {
	core.Camera
	ViewProjection   [6]UniformBufferSection.Matrix4x4
	ProjectionMatrix GeometryMath.Matrix4x4
}

func (camera *Camera) ForceUpdate() {
	for i := range camera.ViewProjection {
		camera.ViewProjection[i].ForceUpdate()
	}
}

func (camera *Camera) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	for i := range camera.ViewProjection {
		camera.ViewProjection[i].SetUniformBuffer(ubo, offset+i*UniformBuffer.Std140_size_mat4)
	}
}

func (camera *Camera) GetSize() int {
	return camera_size_section
}

func (camera *Camera) Update(position GeometryMath.Vector3) {
	camera.ViewProjection[0].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	camera.ViewProjection[1].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{-1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	camera.ViewProjection[2].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, 1.0})))
	camera.ViewProjection[3].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, -1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, -1.0})))
	camera.ViewProjection[4].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 0.0, 1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	camera.ViewProjection[5].Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(GeometryMath.Vector3{0.0, 0.0, -1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.ProjectionMatrix = projection.Decode()
	camera.Camera.SetProjection(projection)
}
