package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

func (light *UBODirectionalLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		UniformBuffer *UniformBuffer.UniformBuffer `yaml:"uniformBuffer"`
	}{
		UniformBuffer: &UniformBuffer.UniformBuffer{
			Size: ubo_size,
			Type: UBO_type,
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return nil
	}

	yamlConfig := struct {
		DirectionalLightSection UBOSection        `yaml:"directionalLight"`
		CameraSection           Camera.UBOSection `yaml:"shadowMap"`
	}{
		DirectionalLightSection: UBOSection{
			DirectionalLight: light.UBOSection.DirectionalLight,
			UniformBuffer:    uboYamlConfig.UniformBuffer,
		},
		CameraSection: Camera.UBOSection{
			Camera:        light.CameraSection.Camera,
			UniformBuffer: uboYamlConfig.UniformBuffer,
			Offset:        offset_camera,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	light.UBOSection = yamlConfig.DirectionalLightSection
	light.CameraSection = yamlConfig.CameraSection

	return nil
}
