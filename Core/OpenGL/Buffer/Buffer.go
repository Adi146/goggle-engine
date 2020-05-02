package Buffer

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type IBuffer interface {
	GetID() uint32
	Destroy()
	Bind()
	Sync()
}

type buffer struct {
	id         uint32
	data       interface{}
	bufferType uint32
	isSync     bool
}

func newBuffer(bufferType uint32, data interface{}) buffer {
	buffer := buffer{
		data:       data,
		bufferType: bufferType,
		isSync:     true,
	}

	gl.GenBuffers(1, &buffer.id)
	buffer.Bind()
	gl.BufferData(buffer.bufferType, Utils.SizeOf(data), Utils.GlPtr(buffer.data), gl.DYNAMIC_DRAW)

	return buffer
}

func (buffer *buffer) GetID() uint32 {
	return buffer.id
}

func (buffer *buffer) Destroy() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer *buffer) Bind() {
	gl.BindBuffer(buffer.bufferType, buffer.id)
}

func (buffer *buffer) Sync() {
	if !buffer.isSync {
		var currentSize int32
		gl.GetNamedBufferParameteriv(buffer.id, gl.BUFFER_SIZE, &currentSize)

		dataSize := Utils.SizeOf(buffer.data)

		if dataSize > int(currentSize) {
			gl.NamedBufferData(buffer.id, dataSize, Utils.GlPtr(buffer.data), gl.DYNAMIC_DRAW)
		} else {
			gl.NamedBufferSubData(buffer.id, 0, dataSize, Utils.GlPtr(buffer.data))
		}

		buffer.isSync = true
	}
}

func (buffer *buffer) Set(data interface{}) {
	buffer.data = data
	buffer.isSync = false
}
