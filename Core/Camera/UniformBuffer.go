package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_projectionMatrix = 0
	offset_viewMatrix       = 64

	ubo_size          = 2 * ubo.Std140_size_mat4
	UBO_type ubo.Type = "camera"
)

type UniformBuffer struct {
	Camera
	ubo.UniformBufferBase
}

func (buff *UniformBuffer) Set(camera Camera) {
	buff.Camera = camera
	buff.ForceUpdate()
}

func (buff *UniformBuffer) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	buff.Camera.SetProjectionMatrix(matrix)
	buff.UpdateData(&matrix[0][0], offset_projectionMatrix, ubo.Std140_size_mat4)
}

func (buff *UniformBuffer) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	buff.Camera.SetViewMatrix(matrix)
	buff.UpdateData(&matrix[0][0], offset_viewMatrix, ubo.Std140_size_mat4)
}

func (buff *UniformBuffer) ForceUpdate() {
	projectionMatrix := buff.ProjectionMatrix
	viewMatrix := buff.ViewMatrix

	buff.UpdateData(&projectionMatrix[0][0], offset_projectionMatrix, ubo.Std140_size_mat4)
	buff.UpdateData(&viewMatrix[0][0], offset_viewMatrix, ubo.Std140_size_mat4)
}
