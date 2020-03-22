package FrameBuffer

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"gopkg.in/yaml.v3"
)

type Type string

type FrameBuffer struct {
	fbo      uint32
	Viewport Viewport
	Type     Type
}

func NewFrameBuffer(viewport Viewport, fboType Type) FrameBuffer {
	buff := FrameBuffer{
		Viewport: viewport,
		Type:     fboType,
	}

	gl.GenFramebuffers(1, &buff.fbo)

	return buff
}

func (buff *FrameBuffer) Destroy() {
	gl.DeleteFramebuffers(1, &buff.fbo)
}

func (buff *FrameBuffer) GetFBO() uint32 {
	return buff.fbo
}

func (buff *FrameBuffer) GetType() Type {
	return buff.Type
}

func (buff *FrameBuffer) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
}

func (buff *FrameBuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.GetFBO())
	buff.Viewport.Bind()
}

func (buff *FrameBuffer) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (buff *FrameBuffer) CheckCompleteness() error {
	if status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER); status != gl.FRAMEBUFFER_COMPLETE {
		return fmt.Errorf("framebuffer is not complete! current status %x", status)
	}

	return nil
}

func (buff *FrameBuffer) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Viewport Viewport `yaml:",inline"`
		Type     Type     `yaml:"type"`
	}{
		Viewport: buff.Viewport,
		Type:     buff.Type,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	*buff = NewFrameBuffer(yamlConfig.Viewport, yamlConfig.Type)
	return nil
}

func GetCurrentFrameBuffer() *FrameBuffer {
	var fbo int32
	gl.GetIntegerv(gl.FRAMEBUFFER_BINDING, &fbo)

	return &FrameBuffer{
		fbo:      uint32(fbo),
		Viewport: GetCurrentViewport(),
	}
}
