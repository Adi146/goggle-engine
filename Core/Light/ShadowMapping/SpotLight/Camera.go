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

func (camera *Camera) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	camera.Set(camera.ProjectionMatrix.Mul(GeometryMath.LookAt(position, position.Add(front), GeometryMath.Vector3{0, 1, 0})))
	camera.Camera.Update(position, front, up)
}

func (camera *Camera) SetProjection(projection GeometryMath.IProjectionConfig) {
	camera.ProjectionMatrix = projection.Decode()
	camera.Camera.SetProjection(projection)
}
