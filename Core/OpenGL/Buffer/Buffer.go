package Buffer

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.3-core/gl"
)

type IBuffer interface {
	GetID() uint32
	Destroy()
	Bind()
	Sync()
}

type Buffer struct {
	id         uint32
	data       interface{}
	bufferType uint32
	usage      uint32
}

func newBuffer(bufferType uint32, data interface{}, usage uint32) Buffer {
	buffer := Buffer{
		data:       data,
		bufferType: bufferType,
		usage:      usage,
	}

	gl.GenBuffers(1, &buffer.id)
	buffer.Bind()
	gl.BufferData(buffer.bufferType, Utils.SizeOf(data), Utils.GlPtr(buffer.data), buffer.usage)

	return buffer
}

func (buffer *Buffer) GetID() uint32 {
	return buffer.id
}

func (buffer *Buffer) Destroy() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer *Buffer) Bind() {
	gl.BindBuffer(buffer.bufferType, buffer.id)
}

func (buffer *Buffer) Sync() {
	if buffer.id == 0 {
		return
	}

	var currentSize int32
	gl.GetNamedBufferParameteriv(buffer.id, gl.BUFFER_SIZE, &currentSize)

	dataSize := Utils.SizeOf(buffer.data)

	if dataSize > int(currentSize) {
		gl.NamedBufferData(buffer.id, dataSize, Utils.GlPtr(buffer.data), buffer.usage)
	} else {
		gl.NamedBufferSubData(buffer.id, 0, dataSize, Utils.GlPtr(buffer.data))
	}
}

func (buffer *Buffer) Set(data interface{}) {
	buffer.data = data
}
