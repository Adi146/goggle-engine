package internal

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_distance = 0

	Size_shadowDistanceSection = 16
)

type ShadowDistanceSection struct {
	Distance      float32
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *ShadowDistanceSection) SetDistance(distance float32) {
	section.Distance = distance
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&distance, section.Offset+offset_distance, UniformBuffer.Std140_size_single)
	}
}

func (section *ShadowDistanceSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *ShadowDistanceSection) GetSize() int {
	return Size_shadowDistanceSection
}

func (section *ShadowDistanceSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		distance := section.Distance

		section.UniformBuffer.UpdateData(&distance, section.Offset+offset_distance, UniformBuffer.Std140_size_single)
	}
}

func (section *ShadowDistanceSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.Distance); err != nil {
		return err
	}

	section.ForceUpdate()

	return nil
}
