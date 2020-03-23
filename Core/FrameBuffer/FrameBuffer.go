package FrameBuffer

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"gopkg.in/yaml.v3"
)

type Type string

type FrameBuffer struct {
	fbo      uint32
	Viewport Viewport
	Type     Type

	ColorAttachments  []IAttachment
	DepthAttachment   IAttachment
	StencilAttachment IAttachment
}

func NewFrameBuffer(viewport Viewport, fboType Type) FrameBuffer {
	var maxColorAttachments int32
	gl.GetIntegerv(gl.MAX_COLOR_ATTACHMENTS, &maxColorAttachments)

	buff := FrameBuffer{
		Viewport:         viewport,
		Type:             fboType,
		ColorAttachments: make([]IAttachment, maxColorAttachments),
	}

	gl.GenFramebuffers(1, &buff.fbo)

	//Fix Error GL_INVALID_OPERATION errors
	buff.Bind()
	buff.Unbind()

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

func (buff *FrameBuffer) AddColorAttachment(attachment IAttachment, index uint32) {
	switch attachment.(type) {
	case Texture.ITexture:
		gl.NamedFramebufferTexture(buff.GetFBO(), gl.COLOR_ATTACHMENT0+index, attachment.GetID(), 0)
	case IAttachment:
		gl.NamedFramebufferRenderbuffer(buff.GetFBO(), gl.COLOR_ATTACHMENT0+index, gl.RENDERBUFFER, attachment.GetID())
	}

	buff.ColorAttachments[index] = attachment
}

func (buff *FrameBuffer) AddDepthAttachment(attachment IAttachment) {
	switch attachment.(type) {
	case Texture.ITexture:
		gl.NamedFramebufferTexture(buff.GetFBO(), gl.DEPTH_ATTACHMENT, attachment.GetID(), 0)
	case IAttachment:
		gl.NamedFramebufferRenderbuffer(buff.GetFBO(), gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, attachment.GetID())
	}

	buff.StencilAttachment = attachment
}

func (buff *FrameBuffer) AddDepthStencilAttachment(attachment IAttachment) {
	switch attachment.(type) {
	case Texture.ITexture:
		gl.NamedFramebufferTexture(buff.GetFBO(), gl.DEPTH_STENCIL_ATTACHMENT, attachment.GetID(), 0)
	case IAttachment:
		gl.NamedFramebufferRenderbuffer(buff.GetFBO(), gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, attachment.GetID())
	}

	buff.DepthAttachment = attachment
	buff.StencilAttachment = attachment
}

func (buff *FrameBuffer) AddStencilAttachment(attachment IAttachment) {
	switch attachment.(type) {
	case Texture.ITexture:
		gl.NamedFramebufferTexture(buff.GetFBO(), gl.STENCIL_ATTACHMENT, attachment.GetID(), 0)
	case IAttachment:
		gl.NamedFramebufferRenderbuffer(buff.GetFBO(), gl.STENCIL_ATTACHMENT, gl.RENDERBUFFER, attachment.GetID())
	}

	buff.DepthAttachment = attachment
}

func (buff *FrameBuffer) Finish() error {
	for _, colorAttachment := range buff.ColorAttachments {
		if colorAttachment != nil {
			return buff.CheckCompleteness()
		}
	}

	gl.NamedFramebufferDrawBuffer(buff.GetFBO(), gl.NONE)
	gl.NamedFramebufferReadBuffer(buff.GetFBO(), gl.NONE)

	return buff.CheckCompleteness()
}

func (buff *FrameBuffer) CheckCompleteness() error {
	if status := gl.CheckNamedFramebufferStatus(buff.GetFBO(), gl.FRAMEBUFFER); status != gl.FRAMEBUFFER_COMPLETE {
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
