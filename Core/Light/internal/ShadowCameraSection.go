package internal

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_viewProjectionMatrix = 0

	Size_shadowCameraSection = 64
)

type ShadowCameraSection struct {
	Camera.Camera
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *ShadowCameraSection) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetProjectionMatrix(matrix)
	if section.UniformBuffer != nil {
		viewProjectionMatrix := *matrix.Mul(&section.Camera.ViewMatrix)
		section.UniformBuffer.UpdateData(&viewProjectionMatrix[0][0], section.Offset+offset_viewProjectionMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *ShadowCameraSection) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	section.Camera.SetViewMatrix(matrix)
	if section.UniformBuffer != nil {
		viewProjectionMatrix := *section.Camera.ProjectionMatrix.Mul(&matrix)
		section.UniformBuffer.UpdateData(&viewProjectionMatrix[0][0], section.Offset+offset_viewProjectionMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *ShadowCameraSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *ShadowCameraSection) GetSize() int {
	return Size_shadowCameraSection
}

func (section *ShadowCameraSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		viewProjectionMatrix := section.Camera.ProjectionMatrix.Mul(&section.Camera.ViewMatrix)

		section.UniformBuffer.UpdateData(&viewProjectionMatrix[0][0], section.Offset+offset_viewProjectionMatrix, UniformBuffer.Std140_size_mat4)
	}
}

func (section *ShadowCameraSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.Camera); err != nil {
		return err
	}

	if section.ViewMatrix == (GeometryMath.Matrix4x4{}) {
		section.ViewMatrix = *GeometryMath.Identity()
	}

	section.ForceUpdate()

	return nil
}
