package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer/MemoryLayout"
	"gopkg.in/yaml.v3"
)

const (
	ubo_binding = 0
)

type UBOCamera struct {
	Camera

	Buffer.UniformBuffer
	Buffer.DynamicBufferData
}

func (camera *UBOCamera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.Camera.Update(position, front, up)
	camera.SetIsSync(false)
}

func (camera *UBOCamera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.Camera.SetProjection(projection)
	camera.SetIsSync(false)
}

func (camera *UBOCamera) GetBufferData() interface{} {
	return &struct {
		projectionMatrix GeometryMath.Matrix4x4
		viewMatrix       GeometryMath.Matrix4x4
		position         MemoryLayout.Std140Vector3
	}{
		projectionMatrix: camera.GetProjectionMatrix(),
		viewMatrix:       camera.GetViewMatrix(),
		position:         MemoryLayout.Std140Vector3{Vector3: camera.GetPosition()},
	}
}

func (camera *UBOCamera) UnmarshalYAML(value *yaml.Node) error {
	camera.UniformBuffer = Buffer.NewUniformBuffer(camera, ubo_binding)

	var yamlConfig GeometryMath.PerspectiveConfig
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	camera.SetProjection(&yamlConfig)

	return nil
}
