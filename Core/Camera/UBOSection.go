package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_projectionMatrix = 0
	offset_viewMatrix       = 64

	ubo_size                    = 2 * UniformBuffer.Std140_size_mat4
	UBO_type UniformBuffer.Type = "camera"
)

type UBOSection struct {
	Camera
	UniformBuffer.UniformBufferBase
}

func (section *UBOSection) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetProjectionMatrix(matrix)
	section.UpdateData(&matrix[0][0], offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
}

func (section *UBOSection) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetViewMatrix(matrix)
	section.UpdateData(&matrix[0][0], offset_viewMatrix, UniformBuffer.Std140_size_mat4)
}

func (section *UBOSection) ForceUpdate() {
	projectionMatrix := section.ProjectionMatrix
	viewMatrix := section.ViewMatrix

	section.UpdateData(&projectionMatrix[0][0], offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
	section.UpdateData(&viewMatrix[0][0], offset_viewMatrix, UniformBuffer.Std140_size_mat4)
}
