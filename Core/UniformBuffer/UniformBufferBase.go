package UniformBuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	Std140_size_single = 4
	Std140_size_vec3   = 12
	Std140_size_mat4   = 64
)

type UniformBufferBase struct {
	ubo     uint32
	Size    int    `yaml:"size"`
	Binding uint32 `yaml:"binding"`
}

func (buff *UniformBufferBase) Init() error {
	gl.GenBuffers(1, &buff.ubo)
	buff.Bind()
	gl.BufferData(gl.UNIFORM_BUFFER, buff.Size, nil, gl.STATIC_DRAW)
	buff.Unbind()

	gl.BindBufferRange(gl.UNIFORM_BUFFER, buff.Binding, buff.ubo, 0, buff.Size)

	return nil
}

func (buff *UniformBufferBase) GetUBO() uint32 {
	return buff.ubo
}

func (buff *UniformBufferBase) Bind() {
	gl.BindBuffer(gl.UNIFORM_BUFFER, buff.ubo)
}

func (buff *UniformBufferBase) Unbind() {
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

func (buff *UniformBufferBase) UpdateData(data interface{}, offset int, size int) {
	gl.NamedBufferSubData(buff.ubo, offset, size, gl.Ptr(data))
}

func (buff *UniformBufferBase) GetBinding() uint32 {
	return buff.Binding
}
