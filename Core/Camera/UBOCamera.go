package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"gopkg.in/yaml.v3"
)

const (
	ubo_binding = 0
)

type UBOCamera struct {
	Camera

	uniformData   *std140CameraData
	uniformBuffer Buffer.UniformBuffer
}

type std140CameraData struct {
	projectionMatrix GeometryMath.Matrix4x4
	viewMatrix       GeometryMath.Matrix4x4
	position         GeometryMath.Vector3
	padding          float32
}

func (camera *UBOCamera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.uniformData.position = position
	camera.uniformData.viewMatrix = GeometryMath.LookAt(position, position.Add(front), up)
	camera.uniformBuffer.Sync()

	camera.Camera.Update(position, front, up)
}

func (camera *UBOCamera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.uniformData.projectionMatrix = projection.Decode()
	camera.uniformBuffer.Sync()

	camera.Camera.SetProjection(projection)
}

func (camera *UBOCamera) UnmarshalYAML(value *yaml.Node) error {
	camera.uniformData = &std140CameraData{}
	camera.uniformBuffer = Buffer.NewUniformBuffer(camera.uniformData, ubo_binding)

	var yamlConfig GeometryMath.PerspectiveConfig
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	camera.SetProjection(&yamlConfig)

	return nil
}
