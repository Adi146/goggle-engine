package Light

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"gopkg.in/yaml.v3"
)

const (
	spotLight_offset_position  = 0
	spotLight_offset_linear    = 12
	spotLight_offset_quadratic = 16
	spotLight_offset_ambient   = 32
	spotLight_offset_diffuse   = 48
	spotLight_offset_specular  = 64
	spotLight_offset_direction = 80
	spotLight_offset_innerCone = 92
	spotLight_offset_outerCone = 96

	spotLight_size_section = 112
	spotLight_ubo_size     = UniformBuffer.Std140_size_single + UniformBuffer.Num_elements*spotLight_size_section
	SpotLight_ubo_type     = "spotLight"
)

type UBOSpotLight struct {
	Position  UniformBufferSection.Vector3
	Linear    UniformBufferSection.Float
	Quadratic UniformBufferSection.Float
	Ambient   UniformBufferSection.Vector3
	Diffuse   UniformBufferSection.Vector3
	Specular  UniformBufferSection.Vector3
	Direction UniformBufferSection.Vector3
	InnerCone UniformBufferSection.Float
	OuterCone UniformBufferSection.Float
}

func (light *UBOSpotLight) ForceUpdate() {
	light.Position.ForceUpdate()
	light.Linear.ForceUpdate()
	light.Quadratic.ForceUpdate()
	light.Ambient.ForceUpdate()
	light.Diffuse.ForceUpdate()
	light.Specular.ForceUpdate()
	light.Direction.ForceUpdate()
	light.InnerCone.ForceUpdate()
	light.OuterCone.ForceUpdate()
}

func (light *UBOSpotLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.Position.SetUniformBuffer(ubo, offset+spotLight_offset_position)
	light.Linear.SetUniformBuffer(ubo, offset+spotLight_offset_linear)
	light.Quadratic.SetUniformBuffer(ubo, offset+spotLight_offset_quadratic)
	light.Ambient.SetUniformBuffer(ubo, offset+spotLight_offset_ambient)
	light.Diffuse.SetUniformBuffer(ubo, offset+spotLight_offset_diffuse)
	light.Specular.SetUniformBuffer(ubo, offset+spotLight_offset_specular)
	light.Direction.SetUniformBuffer(ubo, offset+spotLight_offset_direction)
	light.InnerCone.SetUniformBuffer(ubo, offset+spotLight_offset_innerCone)
	light.OuterCone.SetUniformBuffer(ubo, offset+spotLight_offset_outerCone)
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

	yamlConfig := struct {
		SpotLight `yaml:"spotLight"`
	}{
		SpotLight: SpotLight{
			Position:  light.Position.Get(),
			Linear:    light.Linear.Get(),
			Quadratic: light.Quadratic.Get(),
			Ambient:   light.Ambient.Get(),
			Diffuse:   light.Diffuse.Get(),
			Specular:  light.Specular.Get(),
			Direction: light.Direction.Get(),
			InnerCone: light.InnerCone.Get(),
			OuterCone: light.OuterCone.Get(),
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.Position.Set(yamlConfig.SpotLight.Position)
	light.Linear.Set(yamlConfig.SpotLight.Linear)
	light.Quadratic.Set(yamlConfig.SpotLight.Quadratic)
	light.Ambient.Set(yamlConfig.SpotLight.Ambient)
	light.Diffuse.Set(yamlConfig.SpotLight.Diffuse)
	light.Specular.Set(yamlConfig.SpotLight.Specular)
	light.Direction.Set(yamlConfig.SpotLight.Direction)
	light.InnerCone.Set(yamlConfig.SpotLight.InnerCone)
	light.OuterCone.Set(yamlConfig.SpotLight.OuterCone)

	if _, err := uboYamlConfig.Ptr.AddElement(light); err != nil {
		return err
	}

	return nil
}
