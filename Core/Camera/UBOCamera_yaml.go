package Camera

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

type UBOCamera_yaml struct {
	Section UBOSection `yaml:",inline"`
}

type UBOCamera_uniformBuffer_yaml struct {
	UniformBuffer *UniformBuffer.UniformBuffer `yaml:"uniformBuffer"`
}

func (camera *UBOCamera) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := UBOCamera_uniformBuffer_yaml{
		UniformBuffer: &UniformBuffer.UniformBuffer{
			Size: ubo_size,
			Type: UBO_type,
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return nil
	}

	yamlConfig := UBOCamera_yaml{
		Section: UBOSection{
			Camera:        camera.Camera,
			UniformBuffer: uboYamlConfig.UniformBuffer,
			Offset:        0,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	*camera = (UBOCamera)(yamlConfig.Section)

	return nil
}
