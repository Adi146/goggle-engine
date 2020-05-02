package UniformBuffer

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/go-gl/gl/v4.3-core/gl"
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

func NewUniformBufferBase(size int, binding uint32, uboType Type) (UniformBuffer, error) {
	buff := UniformBuffer{
		Size:    size,
		Binding: binding,
		Type:    uboType,
	}

	var maxUniformBufferSize int32
	gl.GetIntegerv(gl.MAX_UNIFORM_BLOCK_SIZE, &maxUniformBufferSize)
	if int32(size) > maxUniformBufferSize {
		return buff, fmt.Errorf("size of uniform buffer is bigger than max block size(ubo size %d, max block size: %d)", size, maxUniformBufferSize)
	}

	gl.GenBuffers(1, &buff.ubo)
	buff.Bind()
	gl.BufferData(gl.UNIFORM_BUFFER, buff.Size, nil, gl.STATIC_DRAW)
	buff.Unbind()

	gl.BindBufferRange(gl.UNIFORM_BUFFER, buff.Binding, buff.ubo, 0, buff.Size)

	return buff, nil
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

func (buff *UniformBuffer) UpdateData(data interface{}, offset int) {
	gl.NamedBufferSubData(buff.ubo, offset, Utils.SizeOf(data), Utils.GlPtr(data))
}

func (buff *UniformBuffer) GetBinding() uint32 {
	return buff.Binding
}
