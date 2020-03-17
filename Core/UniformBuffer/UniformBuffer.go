package UniformBuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	Std140_size_single = 4
	Std140_size_vec3   = 12
	Std140_size_mat4   = 64
)

type Type string

type UniformBuffer struct {
	ubo     uint32
	Size    int    `yaml:"size"`
	Binding uint32 `yaml:"binding"`
	Type    Type
}

func NewUniformBufferBase(size int, binding uint32, uboType Type) UniformBuffer {
	buff := UniformBuffer{
		Size:    size,
		Binding: binding,
		Type:    uboType,
	}

	gl.GenBuffers(1, &buff.ubo)
	buff.Bind()
	gl.BufferData(gl.UNIFORM_BUFFER, buff.Size, nil, gl.STATIC_DRAW)
	buff.Unbind()

	gl.BindBufferRange(gl.UNIFORM_BUFFER, buff.Binding, buff.ubo, 0, buff.Size)

	return buff
}

func (buff *UniformBuffer) GetUBO() uint32 {
	return buff.ubo
}

func (buff *UniformBuffer) GetType() Type {
	return buff.Type
}

func (buff *UniformBuffer) Bind() {
	gl.BindBuffer(gl.UNIFORM_BUFFER, buff.ubo)
}

func (buff *UniformBuffer) Unbind() {
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

func (buff *UniformBuffer) UpdateData(data interface{}, offset int, size int) {
	ubo := buff.ubo
	gl.NamedBufferSubData(ubo, offset, size, gl.Ptr(data))
}

func (buff *UniformBuffer) GetBinding() uint32 {
	return buff.Binding
}