package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_projectionMatrix = 0
	offset_viewMatrix       = 64
)

type UBOSection struct {
	Camera
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *UBOSection) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetProjectionMatrix(matrix)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&matrix[0][0], section.Offset+offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *UBOSection) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetViewMatrix(matrix)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&matrix[0][0], section.Offset+offset_viewMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *UBOSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		projectionMatrix := section.ProjectionMatrix
		viewMatrix := section.ViewMatrix

		section.UniformBuffer.UpdateData(&projectionMatrix[0][0], section.Offset+offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
		section.UniformBuffer.UpdateData(&viewMatrix[0][0], section.Offset+offset_viewMatrix, UniformBuffer.Std140_size_mat4)
	}
}
