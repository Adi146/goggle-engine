package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	spotLight_offset_position  = 0
	spotLight_offset_color     = 32
	spotLight_offset_direction = 80
	spotLight_offset_cone      = 92

	spotLight_size_section = 112

	spotLight_ubo_size                    = UniformBuffer.Std140_size_single + UniformBuffer.Num_elements*spotLight_size_section
	SpotLight_ubo_type UniformBuffer.Type = "spotLight"
)

type UBOSpotLight struct {
	internal.LightPositionSection
	internal.LightColorSection
	internal.LightDirectionSection
	internal.LightConeSection
}

func (light *UBOSpotLight) ForceUpdate() {
	light.LightPositionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	light.LightDirectionSection.ForceUpdate()
	light.LightConeSection.ForceUpdate()
}

func (light *UBOSpotLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightPositionSection.SetUniformBuffer(ubo, offset+spotLight_offset_position)
	light.LightColorSection.SetUniformBuffer(ubo, offset+spotLight_offset_color)
	light.LightDirectionSection.SetUniformBuffer(ubo, offset+spotLight_offset_direction)
	light.LightConeSection.SetUniformBuffer(ubo, offset+spotLight_offset_cone)
}

func (light *UBOSpotLight) GetSize() int {
	return spotLight_size_section
}
