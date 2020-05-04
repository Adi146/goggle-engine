package Buffer

import (
	"github.com/go-gl/gl/all-core/gl"
)

type DynamicBuffer struct {
	Buffer
	data IDynamicBufferData
}

type IDynamicBufferData interface {
	IsSync() bool
	SetIsSync(val bool)
	GetBufferData() interface{}
}

type DynamicBufferData struct {
	isSync bool
}

func (data *DynamicBufferData) IsSync() bool {
	return data.isSync
}

func (data *DynamicBufferData) SetIsSync(val bool) {
	data.isSync = val
}

func newDynamicBuffer(bufferType uint32, data IDynamicBufferData) DynamicBuffer {
	return DynamicBuffer{
		Buffer: newBuffer(bufferType, data.GetBufferData(), gl.DYNAMIC_DRAW),
		data:   data,
	}
}

func (buffer *DynamicBuffer) Set(data IDynamicBufferData) {
	buffer.data = data
}

func (buffer *DynamicBuffer) Sync() {
	if !buffer.data.IsSync() {
		buffer.Buffer.data = buffer.data.GetBufferData()
		buffer.Buffer.Sync()
		buffer.data.SetIsSync(true)
	}
}
