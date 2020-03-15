package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

type UBOSection_yaml struct {
	Camera        Camera                           `yaml:",inline"`
	UniformBuffer *UniformBuffer.UniformBufferBase `yaml:"uniformBuffer"`
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
	if config.UniformBuffer == nil {
		config.UniformBuffer = &UniformBuffer.UniformBufferBase{
			Size: ubo_size,
			Type: UBO_type,
		}
	}

	if config.Camera.ViewMatrix != (GeometryMath.Matrix4x4{}) {
		config.Camera.ViewMatrix = *GeometryMath.Identity()
	}
}
