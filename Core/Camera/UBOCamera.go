package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
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
	projectionMatrix UniformBufferSection.Matrix4x4
	viewMatrix       UniformBufferSection.Matrix4x4
	position         UniformBufferSection.Vector3

	front GeometryMath.Vector3
	up    GeometryMath.Vector3

	projectionConfig GeometryMath.PerspectiveConfig
	frustum          PlaneFrustum
}

func (camera *UBOCamera) ForceUpdate() {
	camera.projectionMatrix.ForceUpdate()
	camera.viewMatrix.ForceUpdate()
	camera.position.ForceUpdate()
}

func (camera *UBOCamera) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	camera.projectionMatrix.SetUniformBuffer(ubo, offset+camera_offset_projectionMatrix)
	camera.viewMatrix.SetUniformBuffer(ubo, offset+camera_offset_viewMatrix)
	camera.position.SetUniformBuffer(ubo, offset+camera_offset_position)
}

func (camera *UBOCamera) GetSize() int {
	return camera_size_section
}

func (camera *UBOCamera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.position.Set(position)
	camera.viewMatrix.Set(GeometryMath.LookAt(position, position.Add(front), up))

	camera.front = front
	camera.up = up

	camera.frustum.Update(position, front, up)
}

func (camera *UBOCamera) GetViewMatrix() GeometryMath.Matrix4x4 {
	return camera.viewMatrix.Get()
}

func (camera *UBOCamera) GetPosition() GeometryMath.Vector3 {
	return camera.position.Get()
}

func (camera *UBOCamera) GetFront() GeometryMath.Vector3 {
	return camera.front
}

func (camera *UBOCamera) GetUp() GeometryMath.Vector3 {
	return camera.up
}

func (camera *UBOCamera) SetProjection(projection GeometryMath.PerspectiveConfig) {
	camera.projectionMatrix.Set(projection.Decode())
	camera.projectionConfig = projection
	camera.frustum.UpdateProjectionConfig(projection)
}

func (camera *UBOCamera) GetProjection() GeometryMath.PerspectiveConfig {
	return camera.projectionConfig
}

func (camera *UBOCamera) GetProjectionMatrix() GeometryMath.Matrix4x4 {
	return camera.projectionMatrix.Get()
}

func (camera *UBOCamera) GetFrustum() IFrustum {
	return &camera.frustum
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

	var yamlConfig GeometryMath.PerspectiveConfig
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	camera.SetProjection(yamlConfig)
	camera.viewMatrix.Set(GeometryMath.Identity())

	camera.SetUniformBuffer(uboYamlConfig.UniformBuffer, 0)

	return nil
}
