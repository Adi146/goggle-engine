package Camera

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"gopkg.in/yaml.v3"
)

const (
	camera_offset_projectionMatrix = 0
	camera_offset_viewMatrix       = 64
	camera_offset_position         = 128

	camera_size_section = 144
	ubo_size            = camera_size_section
	UBO_type            = "camera"
)

type UBOCamera struct {
	ProjectionMatrix UniformBufferSection.Matrix4x4
	ViewMatrix       UniformBufferSection.Matrix4x4
	Position         UniformBufferSection.Vector3
}

func (camera *UBOCamera) ForceUpdate() {
	camera.ProjectionMatrix.ForceUpdate()
	camera.ViewMatrix.ForceUpdate()
	camera.Position.ForceUpdate()
}

func (camera *UBOCamera) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	camera.ProjectionMatrix.SetUniformBuffer(ubo, offset+camera_offset_projectionMatrix)
	camera.ViewMatrix.SetUniformBuffer(ubo, offset+camera_offset_viewMatrix)
	camera.Position.SetUniformBuffer(ubo, offset+camera_offset_position)
}

func (camera *UBOCamera) GetSize() int {
	return camera_size_section
}

func (camera *UBOCamera) UnmarshalYAML(value *yaml.Node) error {
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

	yamlConfig := Camera{
		ProjectionMatrix: camera.ProjectionMatrix.Get(),
		ViewMatrix:       camera.ViewMatrix.Get(),
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	camera.ProjectionMatrix.Set(yamlConfig.ProjectionMatrix)
	camera.ViewMatrix.Set(yamlConfig.ViewMatrix)

	camera.SetUniformBuffer(uboYamlConfig.UniformBuffer, 0)

	return nil
}
