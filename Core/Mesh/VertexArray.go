package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/go-gl/gl/v4.3-core/gl"
	"unsafe"
)

const (
	index_position = iota
	index_uv
	index_normal
	index_tangent
	index_bitangent

	index_modelMatrix_1
	index_modelMatrix_2
	index_modelMatrix_3
	index_modelMatrix_4
)

type VertexArray uint32

func NewVertexArray(vbo Buffer.Buffer) VertexArray {
	vertex := Vertex{}
	identity := GeometryMath.Identity()

	var vao VertexArray
	gl.GenVertexArrays(1, (*uint32)(&vao))

	vbo.Bind()

	vao.Bind()
	vao.enablePositionAttribute()

	gl.VertexAttribPointer(index_position, int32(len(vertex.Position)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Position)))
	gl.VertexAttribPointer(index_uv, int32(len(vertex.UV)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.UV)))
	gl.VertexAttribPointer(index_normal, int32(len(vertex.Normal)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Normal)))
	gl.VertexAttribPointer(index_tangent, int32(len(vertex.Tangent)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Tangent)))
	gl.VertexAttribPointer(index_bitangent, int32(len(vertex.BiTangent)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.BiTangent)))

	gl.VertexAttrib4fv(index_modelMatrix_1, &identity[0][0])
	gl.VertexAttrib4fv(index_modelMatrix_2, &identity[1][0])
	gl.VertexAttrib4fv(index_modelMatrix_3, &identity[2][0])
	gl.VertexAttrib4fv(index_modelMatrix_4, &identity[3][0])

	return vao
}

func (vao VertexArray) Bind() {
	gl.BindVertexArray(uint32(vao))
}

func (vao VertexArray) Unbind() {
	gl.BindVertexArray(0)
	vao.DisableUVAttribute()
	vao.DisableNormalAttribute()
	vao.DisableTangentAttribute()
	vao.DisableBiTangentAttribute()
}

func (vao VertexArray) enablePositionAttribute() {
	gl.EnableVertexArrayAttrib(uint32(vao), index_position)
}

func (vao VertexArray) disablePositionAttribute() {
	gl.DisableVertexArrayAttrib(uint32(vao), index_position)
}

func (vao VertexArray) EnableUVAttribute() {
	gl.EnableVertexArrayAttrib(uint32(vao), index_uv)
}

func (vao VertexArray) DisableUVAttribute() {
	gl.DisableVertexArrayAttrib(uint32(vao), index_uv)
}

func (vao VertexArray) EnableNormalAttribute() {
	gl.EnableVertexArrayAttrib(uint32(vao), index_normal)
}

func (vao VertexArray) DisableNormalAttribute() {
	gl.DisableVertexArrayAttrib(uint32(vao), index_normal)
}

func (vao VertexArray) EnableTangentAttribute() {
	gl.EnableVertexArrayAttrib(uint32(vao), index_tangent)
}

func (vao VertexArray) DisableTangentAttribute() {
	gl.DisableVertexArrayAttrib(uint32(vao), index_tangent)
}

func (vao VertexArray) EnableBiTangentAttribute() {
	gl.EnableVertexArrayAttrib(uint32(vao), index_bitangent)
}

func (vao VertexArray) DisableBiTangentAttribute() {
	gl.DisableVertexArrayAttrib(uint32(vao), index_bitangent)
}
