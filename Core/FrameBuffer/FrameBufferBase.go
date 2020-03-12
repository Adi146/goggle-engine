package FrameBuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var boundFBO *FrameBufferBase

type FrameBufferBase struct {
	FBO uint32

	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`

	DepthTest bool         `yaml:"depthTest"`
	Culling   CullFunction `yaml:"culling"`
	Blend     bool         `yaml:"blend"`

	previousBuffer *FrameBufferBase
}

func (buff *FrameBufferBase) GetFBO() uint32 {
	return buff.FBO
}

func (buff *FrameBufferBase) GetSize() (int32, int32) {
	return buff.Width, buff.Height
}

func (buff *FrameBufferBase) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
}

func (buff *FrameBufferBase) Bind() {
	buff.rebind()

	buff.previousBuffer = boundFBO
	boundFBO = buff
}

func (buff *FrameBufferBase) rebind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.GetFBO())
	width, height := buff.GetSize()
	gl.Viewport(0, 0, width, height)

	if buff.DepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}

	if buff.Culling.Enabled {
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(buff.Culling.Function)
	} else {
		gl.Disable(gl.CULL_FACE)
	}

	if buff.Blend {
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	} else {
		gl.Disable(gl.BLEND)
	}
}

func (buff *FrameBufferBase) Unbind() {
	if boundFBO == buff && buff.previousBuffer != nil {
		buff.previousBuffer.rebind()
	}
}
