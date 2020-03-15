package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

type UBOSection_yaml struct {
	Camera        Camera                          `yaml:",inline"`
	UniformBuffer UniformBuffer.UniformBufferBase `yaml:"uniformBuffer"`
}

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := UBOSection_yaml{
		Camera:        section.Camera,
		UniformBuffer: section.UniformBufferBase,
	}
	yamlConfig.setDefaults()
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	section.UniformBufferBase = yamlConfig.UniformBuffer
	section.SetProjectionMatrix(yamlConfig.Camera.ProjectionMatrix)
	section.SetViewMatrix(yamlConfig.Camera.ViewMatrix)

	return nil
}

func (config *UBOSection_yaml) setDefaults() {
	if config.UniformBuffer.Size == 0 {
		config.UniformBuffer.Size = ubo_size
	}

	if config.UniformBuffer.Type == "" {
		config.UniformBuffer.Type = UBO_type
	}

	for _, row := range config.Camera.ViewMatrix {
		for _, cell := range row {
			if cell != 0 {
				return
			}
		}
	}

	config.Camera.ViewMatrix = *GeometryMath.Identity()
}
