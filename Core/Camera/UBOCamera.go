package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	ubo_size                    = Size_cameraSection
	UBO_type UniformBuffer.Type = "camera"
)

type UBOCamera CameraSection

func (camera *UBOCamera) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	((*CameraSection)(camera)).SetProjectionMatrix(matrix)
}

func (camera *UBOCamera) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	((*CameraSection)(camera)).SetViewMatrix(matrix)
}

func (camera *UBOCamera) SetPosition(pos GeometryMath.Vector3) {
	((*CameraSection)(camera)).SetPosition(pos)
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

	yamlConfig := struct {
		Section CameraSection `yaml:",inline"`
	}{
		Section: CameraSection{
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
