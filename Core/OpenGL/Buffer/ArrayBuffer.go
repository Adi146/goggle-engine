package Buffer

import "github.com/go-gl/gl/all-core/gl"

func NewStaticArrayBuffer(data interface{}) Buffer {
	return newStaticBuffer(gl.ARRAY_BUFFER, data)
}

func NewDynamicArrayBuffer(data IDynamicBufferData) DynamicBuffer {
	return newDynamicBuffer(gl.ARRAY_BUFFER, data)
}
