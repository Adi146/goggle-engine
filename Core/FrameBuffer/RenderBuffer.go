package FrameBuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type RenderBuffer struct {
	rbo uint32
}

func NewRenderBuffer() *RenderBuffer {
	var buff RenderBuffer

	gl.GenRenderbuffers(1, &buff.rbo)

	return &buff
}

func NewDepth24Stencil8Rbo(width int32, height int32) *RenderBuffer {
	buff := NewRenderBuffer()

	gl.NamedRenderbufferStorage(buff.GetID(), gl.DEPTH24_STENCIL8, width, height)

	return buff
}

func (buff *RenderBuffer) GetID() uint32 {
	return buff.rbo
}
