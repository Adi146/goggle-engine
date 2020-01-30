package FrameBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type FrameBuffer struct {
	fbo uint32

	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`

	DepthTest bool `yaml:"depthTest"`
	Culling   bool `yaml:"culling"`

	shaderProgram Shader.IShaderProgram
}

func (buff *FrameBuffer) GetFBO() uint32 {
	return buff.fbo
}

func (buff *FrameBuffer) GetSize() (int32, int32) {
	return buff.Width, buff.Height
}

func (buff *FrameBuffer) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
}

func (buff *FrameBuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.GetFBO())
	width, height := buff.GetSize()
	gl.Viewport(0, 0, width, height)

	if buff.DepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}

	if buff.Culling {
		gl.Enable(gl.CULL_FACE)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
}
