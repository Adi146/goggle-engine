package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

func (light *UBOPointLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: pointLight_ubo_size,
					Type: PointLight_ubo_type,
				},
			},
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	yamlConfig := struct {
		Light struct {
			PositionSection internal.LightPositionSection `yaml:",inline"`
			ColorSection    internal.LightColorSection    `yaml:",inline"`
		} `yaml:"pointLight"`
	}{
		Light: struct {
			PositionSection internal.LightPositionSection `yaml:",inline"`
			ColorSection    internal.LightColorSection    `yaml:",inline"`
		}{
			PositionSection: internal.LightPositionSection{
				LightPosition: light.LightPositionSection.LightPosition,
				Offset:        pointLight_offset_position,
			},
			ColorSection: internal.LightColorSection{
				LightColor: light.LightColorSection.LightColor,
				Offset:     pointLight_offset_color,
			},
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.LightPositionSection = yamlConfig.Light.PositionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	if err := uboYamlConfig.Ptr.AddElement(light); err != nil {
		return err
	}

	return nil
}
