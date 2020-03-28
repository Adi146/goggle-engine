package VertexBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
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
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, Utils.SizeOf(indices), Utils.GlPtr(indices), gl.STATIC_DRAW)

	return &buff
}

func (buff *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buff.bufferId)
}

func (buff *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
