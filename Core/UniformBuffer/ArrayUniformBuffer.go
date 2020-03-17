package UniformBuffer

import (
	"fmt"
)

const (
	Num_elements = 64

	offset_num_elements = 0
	offset_elements     = 16
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

func (buff *ArrayUniformBuffer) AddElement(elem IUniformBufferSection) error {
	if len(buff.Elements)+1 > Num_elements {
		return fmt.Errorf("buffer exceeded")
	}

	offset := offset_elements
	for _, existingElement := range buff.Elements {
		offset += existingElement.GetSize()
	}

	elem.SetUniformBuffer(buff, offset)
	elem.ForceUpdate()

	buff.Elements = append(buff.Elements, elem)
	buff.UpdateNumElements()

	return nil
}

func (buff *ArrayUniformBuffer) UpdateNumElements() {
	num := int32(len(buff.Elements))
	buff.UpdateData(&num, offset_num_elements, Std140_size_single)
}

func (buff *ArrayUniformBuffer) UpdateData(data interface{}, offset int, size int) {
	if buff.UniformBuffer != nil {
		buff.UniformBuffer.UpdateData(data, offset, size)
	}
}
