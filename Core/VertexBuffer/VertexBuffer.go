package VertexBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

const (
	index_position = iota
	index_uv
	index_normal
	index_tangent
	index_bitangent
)

type VertexBuffer struct {
	bufferId uint32
	vao      uint32
}

func NewVertexBuffer(vertices []Vertex) (*VertexBuffer, error) {
	buff := VertexBuffer{}
	vertex := Vertex{}

	gl.GenVertexArrays(1, &buff.vao)
	gl.BindVertexArray(buff.vao)

	gl.GenBuffers(1, &buff.bufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, buff.bufferId)
	gl.BufferData(gl.ARRAY_BUFFER, Utils.SizeOf(vertices), Utils.GlPtr(vertices), gl.STATIC_DRAW)

	// position data
	buff.enablePositionAttribute()
	gl.VertexAttribPointer(index_position, int32(len(vertex.Position)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Position)))
	// Texture coordinates
	buff.EnableUVAttribute()
	gl.VertexAttribPointer(index_uv, int32(len(vertex.UV)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.UV)))
	// normal vector
	buff.EnableNormalAttribute()
	gl.VertexAttribPointer(index_normal, int32(len(vertex.Normal)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Normal)))
	// tangent
	buff.EnableTangentAttribute()
	gl.VertexAttribPointer(index_tangent, int32(len(vertex.Tangent)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Tangent)))
	// bitangent
	buff.EnableBiTangentAttribute()
	gl.VertexAttribPointer(index_bitangent, int32(len(vertex.BiTangent)), gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.BiTangent)))

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
	buff.DisableUVAttribute()
	buff.DisableNormalAttribute()
	buff.DisableTangentAttribute()
	buff.DisableBiTangentAttribute()
}

func (buff *VertexBuffer) enablePositionAttribute() {
	gl.EnableVertexArrayAttrib(buff.vao, index_position)
}

func (buff *VertexBuffer) disablePositionAttribute() {
	gl.DisableVertexArrayAttrib(buff.vao, index_position)
}

func (buff *VertexBuffer) EnableUVAttribute() {
	gl.EnableVertexArrayAttrib(buff.vao, index_uv)
}

func (buff *VertexBuffer) DisableUVAttribute() {
	gl.DisableVertexArrayAttrib(buff.vao, index_uv)
}

func (buff *VertexBuffer) EnableNormalAttribute() {
	gl.EnableVertexArrayAttrib(buff.vao, index_normal)
}

func (buff *VertexBuffer) DisableNormalAttribute() {
	gl.DisableVertexArrayAttrib(buff.vao, index_normal)
}

func (buff *VertexBuffer) EnableTangentAttribute() {
	gl.EnableVertexArrayAttrib(buff.vao, index_tangent)
}

func (buff *VertexBuffer) DisableTangentAttribute() {
	gl.DisableVertexArrayAttrib(buff.vao, index_tangent)
}

func (buff *VertexBuffer) EnableBiTangentAttribute() {
	gl.EnableVertexArrayAttrib(buff.vao, index_bitangent)
}

func (buff *VertexBuffer) DisableBiTangentAttribute() {
	gl.DisableVertexArrayAttrib(buff.vao, index_bitangent)
}
