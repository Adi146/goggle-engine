package Camera

import (
	"unsafe"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"

	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

type UniformBuffer struct {
	ubo.UniformBufferBase `yaml:",inline"`
}

func (buff *UniformBuffer) Init() error {
	buff.Size = 2 * int(unsafe.Sizeof(Matrix.Matrix4x4{}))

	return buff.UniformBufferBase.Init()
}
