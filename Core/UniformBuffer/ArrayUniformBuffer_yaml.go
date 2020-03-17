package UniformBuffer

import (
	"gopkg.in/yaml.v3"
)

var uboMap = map[string]*ArrayUniformBuffer{}

type ArrayUniformBufferPtr struct {
	*ArrayUniformBuffer
}

func (buff *ArrayUniformBufferPtr) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		ID string `yaml:"id"`
	}
	var configAvailable = false

	if value.Kind == yaml.ScalarNode {
		if err := value.Decode(&yamlConfig.ID); err != nil {
			return err
		}
		configAvailable = false
	} else {
		if err := value.Decode(&yamlConfig); err != nil {
			return err
		}
		configAvailable = true
	}

	ubo, existing := uboMap[yamlConfig.ID]
	if !existing {
		ubo = &ArrayUniformBuffer{}
		uboMap[yamlConfig.ID] = ubo
	}

	if configAvailable && (!existing || ubo.UniformBuffer == nil) {
		uboBase := buff.UniformBuffer
		if err := value.Decode(uboBase); err != nil {
			return err
		}

		ubo.UniformBuffer = uboBase
		ubo.ForceUpdate()
	}

	*buff = ArrayUniformBufferPtr{
		ArrayUniformBuffer: ubo,
	}

	return nil
}
