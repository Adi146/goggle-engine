package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	directionalLight_offset_direction = 0
	directionalLight_offset_color     = 16
	directionalLight_offset_camera    = 64

	directionalLIght_size_section = directionalLight_ubo_size

	directionalLight_ubo_size = 192
	DirectionalLight_ubo_type = "directionalLight"
)

type UBODirectionalLight struct {
	internal.LightDirectionSection
	internal.LightColorSection
	Camera.CameraSection
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.LightDirectionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	light.CameraSection.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightDirectionSection.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.LightColorSection.SetUniformBuffer(ubo, offset+directionalLight_offset_color)
	light.CameraSection.SetUniformBuffer(ubo, offset+directionalLight_offset_camera)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLIght_size_section
}
