package Buffer

import "github.com/go-gl/gl/v4.3-core/gl"

type UniformBuffer struct {
	DynamicBuffer
	binding uint32
}

func NewUniformBuffer(data IDynamicBufferData, binding uint32) UniformBuffer {
	buffer := UniformBuffer{
		DynamicBuffer: newDynamicBuffer(gl.UNIFORM_BUFFER, data),
		binding:       binding,
	}

	gl.BindBufferBase(buffer.bufferType, buffer.binding, buffer.id)

	return buffer
}

func (buffer *UniformBuffer) GetBinding() uint32 {
	return buffer.binding
}
