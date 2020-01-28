package FrameBuffer

import "github.com/Adi146/goggle-engine/Core/Shader"

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
