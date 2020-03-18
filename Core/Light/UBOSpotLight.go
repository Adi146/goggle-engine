package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
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

func (light *UBOSpotLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: spotLight_ubo_size,
					Type: SpotLight_ubo_type,
				},
			},
		},
	}

	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	type Light struct {
		PositionSection  internal.LightPositionSection  `yaml:",inline"`
		ColorSection     internal.LightColorSection     `yaml:",inline"`
		DirectionSection internal.LightDirectionSection `yaml:",inline"`
		ConeSection      internal.LightConeSection      `yaml:",inline"`
	}

	yamlConfig := struct {
		Light `yaml:"spotLight"`
	}{
		Light: Light{
			PositionSection: internal.LightPositionSection{
				LightPosition: light.LightPositionSection.LightPosition,
				Offset:        spotLight_offset_position,
			},
			ColorSection: internal.LightColorSection{
				LightColor: light.LightColorSection.LightColor,
				Offset:     spotLight_offset_color,
			},
			DirectionSection: internal.LightDirectionSection{
				LightDirection: light.LightDirectionSection.LightDirection,
				Offset:         spotLight_offset_direction,
			},
			ConeSection: internal.LightConeSection{
				LightCone: light.LightConeSection.LightCone,
				Offset:    spotLight_offset_cone,
			},
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.LightPositionSection = yamlConfig.Light.PositionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	light.LightDirectionSection = yamlConfig.Light.DirectionSection
	light.LightConeSection = yamlConfig.Light.ConeSection
	if err := uboYamlConfig.Ptr.AddElement(light); err != nil {
		return err
	}

	return nil
}
