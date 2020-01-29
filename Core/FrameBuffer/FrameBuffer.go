package FrameBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type FrameBuffer struct {
	FBO uint32

	Width  int32
	Height int32

	shaderProgram Shader.IShaderProgram
}

func (buff *FrameBuffer) GetFBO() uint32 {
	return buff.FBO
}

func (buff *FrameBuffer) GetSize() (int32, int32) {
	return buff.Width, buff.Height
}

func (buff *FrameBuffer) GetShaderProgram() Shader.IShaderProgram {
	return buff.shaderProgram
}

func (buff *FrameBuffer) SetShaderProgram(shaderProgram Shader.IShaderProgram) {
	buff.shaderProgram = shaderProgram
}

func (buff *FrameBuffer) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
}
