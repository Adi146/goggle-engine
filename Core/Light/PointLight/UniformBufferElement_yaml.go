package PointLight

import (
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

var uboMap = map[string]*UniformBuffer{}

type yamlConfig struct {
	ubo.YamlConfig `yaml:",inline"`

	ID string `yaml:"id"`
}

func (elem *UniformBufferElement) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := yamlConfig{
		YamlConfig: ubo.YamlConfig{
			Size: ubo_size,
			Type: UBO_type,
		},
	}
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
		tmpBuff, err := yamlConfig.Create()
		if err != nil {
			return err
		}
		ubo.UniformBufferBase = tmpBuff
		ubo.ForceUpdate()
	}

	return ubo.AddElement(elem)
}
