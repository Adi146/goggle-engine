package Shadow

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/go-gl/gl/all-core/gl"
)

type DepthBuffer struct {
	FrameBuffer.FrameBuffer
	texture Model.Texture

	Width  int32
	Height int32

	activeShaderProgram Shader.IShaderProgram
}

func (buff *DepthBuffer) Init() error {
	gl.GenFramebuffers(1, &buff.FBO)

	gl.BindTexture(gl.TEXTURE_2D, buff.texture.TextureID)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT, buff.Width, buff.Height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.FBO)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.TEXTURE_2D, buff.texture.TextureID, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return nil
}

func (buff *DepthBuffer) Destroy() {
	gl.DeleteFramebuffers(1, &buff.FBO)
}

/*func (depthMap *DepthBuffer) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT)
}*/

func (buff *DepthBuffer) GetSize() (int32, int32) {
	return buff.Width, buff.Height
}

func (buff *DepthBuffer) GetFBO() uint32 {
	return buff.FBO
}

func (buff *DepthBuffer) GetTextures() []*Model.Texture {
	return []*Model.Texture{
		&buff.texture,
	}
}
