package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

type UBOSection_yaml struct {
	DirectionalLight DirectionalLight  `yaml:"directionalLight"`
	CameraSection    Camera.UBOSection `yaml:"shadowMap"`
}

type UBOSection_uniformBuffer_yaml struct {
	UniformBuffer *ubo.UniformBufferBase `yaml:"uniformBuffer"`
}

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := UBOSection_uniformBuffer_yaml{
		UniformBuffer: section.UniformBufferBase,
	}
	uboYamlConfig.setDefaults()
	if err := value.Decode(&uboYamlConfig); err != nil {
		return nil
	}
	section.UniformBufferBase = uboYamlConfig.UniformBuffer

	yamlConfig := UBOSection_yaml{
		DirectionalLight: section.DirectionalLight,
		CameraSection: Camera.UBOSection{
			UniformBufferBase: section.UniformBufferBase,
			Offset:            offset_camera,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	section.SetDirection(yamlConfig.DirectionalLight.Direction)
	section.SetAmbient(yamlConfig.DirectionalLight.Ambient)
	section.SetDiffuse(yamlConfig.DirectionalLight.Diffuse)
	section.SetSpecular(yamlConfig.DirectionalLight.Specular)
	section.UBOSection = yamlConfig.CameraSection

	return nil
}

func (config *UBOSection_uniformBuffer_yaml) setDefaults() {
	if config.UniformBuffer == nil {
		config.UniformBuffer = &ubo.UniformBufferBase{
			Size: ubo_size,
			Type: UBO_type,
		}
	}
}
