package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_projectionMatrix = 0
	offset_viewMatrix       = 64

	Size_cameraSection = 128
)

type CameraSection struct {
	Camera
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *CameraSection) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetProjectionMatrix(matrix)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&matrix[0][0], section.Offset+offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *CameraSection) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetViewMatrix(matrix)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&matrix[0][0], section.Offset+offset_viewMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *CameraSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		projectionMatrix := section.ProjectionMatrix
		viewMatrix := section.ViewMatrix

		section.UniformBuffer.UpdateData(&projectionMatrix[0][0], section.Offset+offset_projectionMatrix, UniformBuffer.Std140_size_mat4)
		section.UniformBuffer.UpdateData(&viewMatrix[0][0], section.Offset+offset_viewMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *CameraSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *CameraSection) GetSize() int {
	return Size_cameraSection
}

func (section *CameraSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.Camera); err != nil {
		return err
	}

	if section.ViewMatrix == (GeometryMath.Matrix4x4{}) {
		section.ViewMatrix = *GeometryMath.Identity()
	}

	section.ForceUpdate()

	return nil
}
