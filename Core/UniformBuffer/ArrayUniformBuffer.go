package UniformBuffer

import (
	"fmt"
)

const (
	Num_elements = 32

	offset_num_elements                = 0
	ArrayUniformBuffer_offset_elements = 16
)

type ArrayUniformBuffer struct {
	*UniformBuffer
	Elements []IUniformBufferSection
}

func (buff *ArrayUniformBuffer) ForceUpdate() {
	buff.UpdateNumElements()
	for _, elem := range buff.Elements {
		elem.ForceUpdate()
	}
}

func (buff *ArrayUniformBuffer) AddElement(elem IUniformBufferSection) (int, error) {
	index := len(buff.Elements)
	if index+1 > Num_elements {
		return index, fmt.Errorf("buffer exceeded")
	}

	offset := ArrayUniformBuffer_offset_elements
	for _, existingElement := range buff.Elements {
		offset += existingElement.GetSize()
	}

	elem.SetUniformBuffer(buff, offset)
	elem.ForceUpdate()

	buff.Elements = append(buff.Elements, elem)
	buff.UpdateNumElements()

	return index, nil
}

func (buff *ArrayUniformBuffer) UpdateNumElements() {
	num := int32(len(buff.Elements))
	buff.UpdateData(&num, offset_num_elements)
}

func (buff *ArrayUniformBuffer) UpdateData(data interface{}, offset int) {
	if buff.UniformBuffer != nil {
		buff.UniformBuffer.UpdateData(data, offset)
	}
}
