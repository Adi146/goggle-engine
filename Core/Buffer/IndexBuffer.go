package Buffer

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

type IndexBuffer struct {
	bufferId uint32
	Length   int32
}

func NewIndexBuffer(indices []uint32) *IndexBuffer {
	buff := IndexBuffer{
		Length: int32(len(indices)),
	}

	gl.GenBuffers(1, &buff.bufferId)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buff.bufferId)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*(int)(unsafe.Sizeof(indices[0])), gl.Ptr(indices), gl.STATIC_DRAW)

	return &buff
}

func (buff *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buff.bufferId)
}

func (buff *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
