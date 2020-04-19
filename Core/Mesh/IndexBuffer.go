package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type IndexBuffer struct {
	bufferId uint32
	Length   int32

	RestartIndex uint32
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
	if buff.RestartIndex != 0 {
		gl.Enable(gl.PRIMITIVE_RESTART)
		gl.PrimitiveRestartIndex(buff.RestartIndex)
	}
}

func (buff *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	if buff.RestartIndex != 0 {
		gl.Disable(gl.PRIMITIVE_RESTART)
	}
}
