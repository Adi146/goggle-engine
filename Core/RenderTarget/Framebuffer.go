package RenderTarget

import "github.com/go-gl/gl/all-core/gl"

type Framebuffer struct {
	fbo      uint32
	textures [2]uint32

	Width  int32
	Height int32
}

func (buff *Framebuffer) Init() {
	gl.GenFramebuffers(1, &buff.fbo)
	gl.GenTextures(2, &buff.textures[0])

	gl.BindTexture(gl.TEXTURE_2D, buff.textures[0])
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, buff.Width, buff.Height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.BindTexture(gl.TEXTURE_2D, buff.textures[1])
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH24_STENCIL8, buff.Width, buff.Height, 0, gl.DEPTH_STENCIL, gl.UNSIGNED_INT_24_8, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	buff.Bind()
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, buff.textures[0], 0)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.TEXTURE_2D, buff.textures[1], 0)
	buff.Unbind()
}

func (buff *Framebuffer) Destroy() {
	gl.DeleteFramebuffers(1, &buff.fbo)
}

func (buff *Framebuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, buff.fbo)
}

func (buff *Framebuffer) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
