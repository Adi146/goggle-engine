package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	directionalLight_offset_direction = 0
	directionalLight_offset_color     = 16
	directionalLight_offset_camera    = 64

	directionalLIght_size_section = directionalLight_ubo_size

	directionalLight_ubo_size = 192
	DirectionalLight_ubo_type = "directionalLight"
)

type UBODirectionalLight struct {
	internal.LightDirectionSection
	internal.LightColorSection
	Camera.CameraSection
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.LightDirectionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	light.CameraSection.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightDirectionSection.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.LightColorSection.SetUniformBuffer(ubo, offset+directionalLight_offset_color)
	light.CameraSection.SetUniformBuffer(ubo, offset+directionalLight_offset_camera)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLIght_size_section
}

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

	type Light struct {
		DirectionSection internal.LightDirectionSection `yaml:",inline"`
		ColorSection     internal.LightColorSection     `yaml:",inline"`
	}

	yamlConfig := struct {
		Light         `yaml:"directionalLight"`
		CameraSection Camera.CameraSection `yaml:"shadowMap"`
	}{
		Light: Light{
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
