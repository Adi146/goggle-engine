package SpotLight

import (
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
					Size: ubo_size,
					Type: UBO_type,
				},
			},
		},
	}

	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	yamlConfig := struct {
		PointLightSection UBOSection `yaml:"spotLight"`
	}{
		PointLightSection: UBOSection{
			SpotLight: light.UBOSection.SpotLight,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.UBOSection = yamlConfig.PointLightSection
	if err := uboYamlConfig.Ptr.AddElement(&light.UBOSection); err != nil {
		return err
	}

	return nil
}
