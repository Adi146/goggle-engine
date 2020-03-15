package SpotLight

import (
	uboCore "github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

var uboMap = map[string]*UniformBuffer{}

type yamlConfig struct {
	ID string `yaml:"id"`
}

func (elem *UniformBufferElement) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig yamlConfig
	var configAvailable = false

	if value.Kind == yaml.ScalarNode {
		value.Decode(&yamlConfig.ID)
		configAvailable = false
	} else {
		value.Decode(&yamlConfig)
		configAvailable = true
	}

	ubo, existing := uboMap[yamlConfig.ID]
	if !existing {
		ubo = &UniformBuffer{}
		uboMap[yamlConfig.ID] = ubo
	}

	if configAvailable && (!existing || ubo.Size == 0) {
		uboBase := uboCore.UniformBufferBase{
			Size: ubo_size,
			Type: UBO_type,
		}

		if err := value.Decode(&uboBase); err != nil {
			return err
		}

		ubo.UniformBufferBase = uboBase
		ubo.ForceUpdate()
	}

	return ubo.AddElement(elem)
}
