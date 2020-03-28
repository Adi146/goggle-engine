package VertexBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexBuffer struct {
	bufferId uint32
	vao      uint32
}

func NewVertexBuffer(vertices interface{}, vertexBufferAttribFunc func()) (*VertexBuffer, error) {
	buff := VertexBuffer{}

	gl.GenVertexArrays(1, &buff.vao)
	gl.BindVertexArray(buff.vao)

	gl.GenBuffers(1, &buff.bufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, buff.bufferId)
	gl.BufferData(gl.ARRAY_BUFFER, Utils.SizeOf(vertices), Utils.GlPtr(vertices), gl.STATIC_DRAW)

	vertexBufferAttribFunc()

	gl.BindVertexArray(0)

	return &buff, nil
}

func (buff *VertexBuffer) Destroy() {
	gl.DeleteBuffers(1, &buff.bufferId)
}

func (buff *VertexBuffer) Bind() {
	gl.BindVertexArray(buff.vao)
}

func (buff *VertexBuffer) Unbind() {
	gl.BindVertexArray(0)
}
