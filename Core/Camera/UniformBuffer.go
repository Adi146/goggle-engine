package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

type UniformBuffer struct {
	Camera
	ubo.UniformBufferBase `yaml:",inline"`
}

func (buff *UniformBuffer) Init() error {
	buff.Size = 2 * 64

	return buff.UniformBufferBase.Init()
}

func (buff *UniformBuffer) Set(camera Camera) {
	buff.Camera = camera
	buff.ForceUpdate()
}

func (buff *UniformBuffer) SetProjectionMatrix(matrix Matrix.Matrix4x4) {
	buff.Camera.SetProjectionMatrix(matrix)
	buff.UpdateData(&matrix[0][0], 0, 64)
}

func (buff *UniformBuffer) SetViewMatrix(matrix Matrix.Matrix4x4) {
	buff.Camera.SetViewMatrix(matrix)
	buff.UpdateData(&matrix[0][0], 64, 64)
}

func (buff *UniformBuffer) ForceUpdate() {
	buff.UpdateData(&buff.ProjectionMatrix[0][0], 0, 64)
	buff.UpdateData(&buff.ViewMatrix[0][0], 64, 0)
}
