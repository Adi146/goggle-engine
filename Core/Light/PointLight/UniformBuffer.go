package PointLight

import (
	"fmt"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	num_elements = 64

	num_lights_offset = 0
)

type UniformBuffer struct {
	ubo.UniformBufferBase `yaml:",inline"`
	Elements              []*UniformBufferElement `yaml:"lights"`
}

func (buff *UniformBuffer) Init() error {
	buff.Size = ubo.Std140_size_single + num_elements*element_size

	err := buff.UniformBufferBase.Init()
	if err != nil {
		return err
	}

	buff.ForceUpdate()

	return nil
}

func (buff *UniformBuffer) ForceUpdate() {
	buff.UpdateNumLights()
	for _, elem := range buff.Elements {
		elem.ForceUpdate()
	}
}

func (buff *UniformBuffer) GetNewElement() (*UniformBufferElement, error) {
	nextIndex := len(buff.Elements)

	if nextIndex+1 > num_elements {
		return nil, fmt.Errorf("buffer exceeded")
	}

	elem := &UniformBufferElement{
		ubo:          buff,
		ElementIndex: nextIndex,
	}

	buff.Elements = append(buff.Elements, elem)
	buff.UpdateNumLights()

	return elem, nil
}

func (buff *UniformBuffer) UpdateNumLights() {
	num := int32(len(buff.Elements))
	buff.UpdateData(&num, num_lights_offset, ubo.Std140_size_single)
}
