package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_position  = 0
	offset_linear    = 12
	offset_quadratic = 16

	Size_lightPositionSection = 32
)

type LightPositionSection struct {
	LightPosition
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *LightPositionSection) SetPosition(pos GeometryMath.Vector3) {
	section.LightPosition.SetPosition(pos)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightPositionSection) SetLinear(val float32) {
	section.LightPosition.SetLinear(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_linear, UniformBuffer.Std140_size_single)
	}
}

func (section *LightPositionSection) SetQuadratic(val float32) {
	section.LightPosition.SetQuadratic(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_quadratic, UniformBuffer.Std140_size_single)
	}
}

func (section *LightPositionSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		pos := section.Position
		linear := section.Linear
		quadratic := section.Quadratic

		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&linear, section.Offset+offset_linear, UniformBuffer.Std140_size_single)
		section.UniformBuffer.UpdateData(&quadratic, section.Offset+offset_quadratic, UniformBuffer.Std140_size_single)
	}
}

func (section *LightPositionSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *LightPositionSection) GetSize() int {
	return Size_lightPositionSection
}

func (section *LightPositionSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightPosition); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
