package FrameBuffer

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type FrameBufferBase struct {
	fbo uint32

	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`

	DepthTest bool `yaml:"depthTest"`
	Culling   bool `yaml:"culling"`
}

func (buff *FrameBufferBase) GetFBO() uint32 {
	return buff.fbo
}

func (buff *FrameBufferBase) GetSize() (int32, int32) {
	return buff.Width, buff.Height
}

func (buff *FrameBufferBase) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
}

func (buff *FrameBufferBase) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.GetFBO())
	fmt.Println(gl.GetError())
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
