package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

func (light *UBODirectionalLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		UniformBuffer *UniformBuffer.UniformBuffer `yaml:"uniformBuffer"`
	}{
		UniformBuffer: &UniformBuffer.UniformBuffer{
			Size: directionalLight_ubo_size,
			Type: DirectionalLight_ubo_type,
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	yamlConfig := struct {
		Light struct {
			DirectionSection internal.LightDirectionSection `yaml:",inline"`
			ColorSection     internal.LightColorSection     `yaml:",inline"`
		} `yaml:"directionalLight"`
		CameraSection Camera.CameraSection `yaml:"shadowMap"`
	}{
		Light: struct {
			DirectionSection internal.LightDirectionSection `yaml:",inline"`
			ColorSection     internal.LightColorSection     `yaml:",inline"`
		}{
			DirectionSection: internal.LightDirectionSection{
				LightDirection: light.LightDirection,
				UniformBuffer:  uboYamlConfig.UniformBuffer,
				Offset:         directionalLight_offset_direction,
			},
			ColorSection: internal.LightColorSection{
				LightColor:    internal.LightColor{},
				UniformBuffer: uboYamlConfig.UniformBuffer,
				Offset:        directionalLight_offset_color,
			},
		},
		CameraSection: Camera.CameraSection{
			Camera:        light.Camera,
			UniformBuffer: uboYamlConfig.UniformBuffer,
			Offset:        directionalLight_offset_camera,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	light.LightDirectionSection = yamlConfig.Light.DirectionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	light.CameraSection = yamlConfig.CameraSection

	return nil
}
