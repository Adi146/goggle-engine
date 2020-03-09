package ShadowMapShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShadowMapBuffer struct {
	FrameBuffer.FrameBufferBase `yaml:",inline"`
	ShadowMap                   *Texture.Texture
}

func (buff *ShadowMapBuffer) Init() error {
	var err error
	if err = gl.Init(); err != nil {
		return err
	}

	shadowMap, err := NewShadowMap(buff.Width, buff.Height)
	if err != nil {
		return err
	}
	buff.ShadowMap = shadowMap

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
		buff.ShadowMap,
	}
}
