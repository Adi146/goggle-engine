package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ArrayBuffer uint32

func NewVertexBuffer(vertices []Vertex) ArrayBuffer {
	var vbo ArrayBuffer

	gl.GenBuffers(1, (*uint32)(&vbo))
	vbo.Bind()
	defer vbo.Unbind()
	gl.BufferData(gl.ARRAY_BUFFER, Utils.SizeOf(vertices), Utils.GlPtr(vertices), gl.STATIC_DRAW)

	return vbo
}

func NewMatrixBuffer(matrices []GeometryMath.Matrix4x4) ArrayBuffer {
	var vbo ArrayBuffer

	gl.GenBuffers(1, (*uint32)(&vbo))
	vbo.Bind()
	defer vbo.Unbind()
	gl.BufferData(gl.ARRAY_BUFFER, Utils.SizeOf(matrices), Utils.GlPtr(matrices), gl.STREAM_DRAW)

	return vbo
}

func (buffer ArrayBuffer) Destroy() {
	gl.DeleteBuffers(1, (*uint32)(&buffer))
}

func (buffer ArrayBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(buffer))
}

func (buffer ArrayBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
