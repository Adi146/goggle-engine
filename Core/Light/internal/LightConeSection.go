package internal

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_innerCone = 0
	offset_outerCone = 4

	Size_lightConeSection = 16
)

type LightConeSection struct {
	LightCone
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *LightConeSection) SetInnerCone(val float32) {
	section.LightCone.SetInnerCone(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_innerCone, UniformBuffer.Std140_size_single)
	}
}

func (section *LightConeSection) SetOuterCone(val float32) {
	section.LightCone.SetOuterCone(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_outerCone, UniformBuffer.Std140_size_single)
	}
}

func (section *LightConeSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		innerCone := section.InnerCone
		outerCone := section.OuterCone

		section.UniformBuffer.UpdateData(&innerCone, section.Offset+offset_innerCone, UniformBuffer.Std140_size_single)
		section.UniformBuffer.UpdateData(&outerCone, section.Offset+offset_outerCone, UniformBuffer.Std140_size_single)
	}
}

func (section *LightConeSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *LightConeSection) GetSize() int {
	return Size_lightConeSection
}

func (section *LightConeSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightCone); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
