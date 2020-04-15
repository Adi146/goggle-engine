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
	gl.BufferData(gl.ARRAY_BUFFER, Utils.SizeOf(matrices), Utils.GlPtr(matrices), gl.DYNAMIC_DRAW)

	return vbo
}

func (buffer ArrayBuffer) UpdateData(data interface{}, offset int) {
	gl.NamedBufferSubData(uint32(buffer), offset, Utils.SizeOf(data), Utils.GlPtr(data))
}

func (buffer ArrayBuffer) IncreaseSize(size int) {
	var currentSize int32
	gl.GetNamedBufferParameteriv(uint32(buffer), gl.BUFFER_SIZE, &currentSize)

	tmp := make([]uint8, currentSize)
	gl.GetNamedBufferSubData(uint32(buffer), 0, int(currentSize), Utils.GlPtr(tmp))

	gl.NamedBufferData(uint32(buffer), int(currentSize)+size, Utils.GlPtr(tmp), gl.DYNAMIC_DRAW)
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
