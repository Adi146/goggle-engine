package UniformBuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	Std140_size_single = 4
	Std140_size_vec3 = 16
	Std140_size_mat4 = 64
)

type UniformBufferBase struct {
	ubo   uint32
	Size  int    `yaml:"size"`
	Index uint32 `yaml:"index"`
}

func (buff *UniformBufferBase) Init() error {
	gl.GenBuffers(1, &buff.ubo)
	buff.Bind()
	gl.BufferData(gl.UNIFORM_BUFFER, buff.Size, nil, gl.STATIC_DRAW)
	buff.Unbind()

	gl.BindBufferRange(gl.UNIFORM_BUFFER, buff.Index, buff.ubo, 0, buff.Size)

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
	buff.Bind()
	gl.BufferSubData(gl.UNIFORM_BUFFER, offset, size, gl.Ptr(data))
	buff.Unbind()
}

func (buff *UniformBufferBase) GetIndex() uint32 {
	return buff.Index
}
