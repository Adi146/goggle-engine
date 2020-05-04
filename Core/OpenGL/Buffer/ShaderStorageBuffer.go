package Buffer

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

type ShaderStorageBuffer struct {
	DynamicBuffer
	binding uint32
}

func NewShaderStorageBuffer(data IDynamicBufferData, binding uint32) ShaderStorageBuffer {
	buffer := ShaderStorageBuffer{
		DynamicBuffer: newDynamicBuffer(gl.SHADER_STORAGE_BUFFER, data),
		binding:       binding,
	}

	gl.BindBufferBase(buffer.bufferType, buffer.binding, buffer.id)

	return buffer
}

func (buffer *ShaderStorageBuffer) GetBinding() uint32 {
	return buffer.binding
}
