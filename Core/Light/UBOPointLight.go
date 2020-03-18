package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
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

	type Light struct {
		PositionSection internal.LightPositionSection `yaml:",inline"`
		ColorSection    internal.LightColorSection    `yaml:",inline"`
	}

	yamlConfig := struct {
		Light `yaml:"pointLight"`
	}{
		Light: Light {
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
