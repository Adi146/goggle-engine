package FrameBuffer

import (
	"fmt"

	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type OffScreenBuffer struct {
	FrameBufferBase `yaml:",inline"`
	ColorTexture    Texture.Texture

	rbo uint32
}

func (buff *OffScreenBuffer) Init() error {
	var err error
	if err = gl.Init(); err != nil {
		return err
	}

	gl.GenTextures(1, &buff.ColorTexture.TextureID)
	gl.BindTexture(gl.TEXTURE_2D, buff.ColorTexture.TextureID)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, buff.Width, buff.Height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.GenRenderbuffers(1, &buff.rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buff.rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, buff.Width, buff.Height)
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)

	gl.GenFramebuffers(1, &buff.fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.fbo)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, buff.ColorTexture.TextureID, 0)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, buff.rbo)

	if status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER); status != gl.FRAMEBUFFER_COMPLETE {
		err = fmt.Errorf("framebuffer is not complete! current status %x", status)
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return err
}

func (buff *OffScreenBuffer) Destroy() {
	gl.DeleteFramebuffers(1, &buff.fbo)
	gl.DeleteRenderbuffers(1, &buff.rbo)
	//buff.ColorTexture.Destroy()
}

func (buff *OffScreenBuffer) GetTextures() []*Texture.Texture {
	return []*Texture.Texture{
		&buff.ColorTexture,
	}
}
