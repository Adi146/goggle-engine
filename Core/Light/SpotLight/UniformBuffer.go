package SpotLight

import (
	"fmt"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	num_elements = 64

	offset_num_lights = 0

	ubo_size          = ubo.Std140_size_single + num_elements*element_size
	UBO_type ubo.Type = "spotLight"
)

type UniformBuffer struct {
	ubo.UniformBuffer
	Elements []*UniformBufferElement
}

func (buff *UniformBuffer) ForceUpdate() {
	buff.UpdateNumLights()
	for _, elem := range buff.Elements {
		elem.ForceUpdate()
	}
}

func (buff *UniformBuffer) AddElement(elem *UniformBufferElement) error {
	nextIndex := len(buff.Elements)

	if nextIndex+1 > num_elements {
		return fmt.Errorf("buffer exceeded")
	}

	elem.ubo = buff
	elem.ElementIndex = nextIndex

	buff.Elements = append(buff.Elements, elem)
	buff.UpdateNumLights()

	return nil
}

func (buff *UniformBuffer) UpdateNumLights() {
	num := int32(len(buff.Elements))
	buff.UpdateData(&num, offset_num_lights, ubo.Std140_size_single)
}
