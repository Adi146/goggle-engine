package ShadowMapShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	ShadowMap Texture.Type = "shadowMap"
)

type ShadowMapBuffer struct {
	FrameBuffer.FrameBufferBase `yaml:",inline"`
	ShadowMap                   Texture.Texture
}

func (buff *ShadowMapBuffer) Init() error {
	buff.ShadowMap.Type = ShadowMap
	buff.ShadowMap.Target = gl.TEXTURE_2D

	var err error
	if err = gl.Init(); err != nil {
		return err
	}

	gl.GenTextures(1, &buff.ShadowMap.ID)
	if err := buff.ShadowMap.Bind(); err != nil {
		return err
	}
	gl.TexImage2D(buff.ShadowMap.Target, 0, gl.DEPTH_COMPONENT24, buff.Width, buff.Height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(buff.ShadowMap.Target, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(buff.ShadowMap.Target, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(buff.ShadowMap.Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(buff.ShadowMap.Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	buff.ShadowMap.Unbind()

	gl.GenFramebuffers(1, &buff.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.FBO)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, buff.ShadowMap.Target, buff.ShadowMap.ID, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)

	if status := gl.CheckFramebufferStatus(gl.FRAMEBUFFER); status != gl.FRAMEBUFFER_COMPLETE {
		err = fmt.Errorf("ShadowMap Framebuffer is not complete! current status %x", status)
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return err
}

func (buff *ShadowMapBuffer) Destroy() {
	gl.DeleteFramebuffers(1, &buff.FBO)
}

func (buff *ShadowMapBuffer) GetTextures() []*Texture.Texture {
	return []*Texture.Texture{
		&buff.ShadowMap,
	}
}
