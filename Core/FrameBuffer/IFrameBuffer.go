package FrameBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type IFrameBuffer interface {
	Init() error
	Destroy()

	GetSize() (int32, int32)
	GetFBO() uint32

	GetShaderProgram() Shader.IShaderProgram
	SetShaderProgram(shaderProgram Shader.IShaderProgram)

	Clear()

	GetTextures() []*Model.Texture
}
