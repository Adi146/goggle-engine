package Buffer

import "github.com/go-gl/gl/v4.1-core/gl"

type ArrayBuffer struct {
	buffer
}

func NewArrayBuffer(data interface{}) ArrayBuffer {
	return ArrayBuffer{
		buffer: newBuffer(gl.ARRAY_BUFFER, data),
	}
}
