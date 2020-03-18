package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

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

	yamlConfig := struct {
		Light struct {
			PositionSection  internal.LightPositionSection  `yaml:",inline"`
			ColorSection     internal.LightColorSection     `yaml:",inline"`
			DirectionSection internal.LightDirectionSection `yaml:",inline"`
			ConeSection      internal.LightConeSection      `yaml:",inline"`
		} `yaml:"spotLight"`
	}{
		Light: struct {
			PositionSection  internal.LightPositionSection  `yaml:",inline"`
			ColorSection     internal.LightColorSection     `yaml:",inline"`
			DirectionSection internal.LightDirectionSection `yaml:",inline"`
			ConeSection      internal.LightConeSection      `yaml:",inline"`
		}{
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
