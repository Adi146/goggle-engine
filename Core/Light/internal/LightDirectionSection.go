package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_direction = 0

	Size_lightDirectionSection = 16
)

type LightDirectionSection struct {
	LightDirection
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *LightDirectionSection) SetDirection(direction GeometryMath.Vector3) {
	section.LightDirection.SetDirection(direction)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&direction[0], section.Offset+offset_direction, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightDirectionSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		direction := section.LightDirection.Direction

		section.UniformBuffer.UpdateData(&direction[0], section.Offset+offset_direction, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightDirectionSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *LightDirectionSection) GetSize() int {
	return Size_lightDirectionSection
}
