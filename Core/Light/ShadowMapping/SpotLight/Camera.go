package SpotLight

import (
	core "github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
)

type Camera struct {
	core.Camera
	UniformBufferSection.Matrix4x4
	ProjectionMatrix GeometryMath.Matrix4x4
}

func (camera *Camera) Update(position GeometryMath.Vector3, direction GeometryMath.Vector3) {
	camera.Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(direction), GeometryMath.Vector3{0, 1, 0})))
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.ProjectionMatrix = projection.Decode()
	camera.Camera.SetProjection(projection)
}
