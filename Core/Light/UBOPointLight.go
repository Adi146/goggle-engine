package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	pointLight_offset_position = 0
	pointLight_offset_color    = 32

	pointLight_size_section = 80

	pointLight_ubo_size                    = UniformBuffer.Std140_size_single + UniformBuffer.Num_elements*pointLight_size_section
	PointLight_ubo_type UniformBuffer.Type = "pointLight"
)

type UBOPointLight struct {
	internal.LightPositionSection
	internal.LightColorSection
}

func (light *UBOPointLight) ForceUpdate() {
	light.LightPositionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
}

func (light *UBOPointLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightPositionSection.SetUniformBuffer(ubo, offset+pointLight_offset_position)
	light.LightColorSection.SetUniformBuffer(ubo, offset+pointLight_offset_color)
}

func (light *UBOPointLight) GetSize() int {
	return pointLight_size_section
}
