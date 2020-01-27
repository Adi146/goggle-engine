package Shadow

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/go-gl/gl/all-core/gl"
)

type DepthMap struct {
	fbo     uint32
	texture uint32

	Width  int32
	Height int32

	activeShaderProgram Shader.IShaderProgram
}

func (depthMap *DepthMap) Init() {
	gl.GenFramebuffers(1, &depthMap.fbo)

	gl.BindTexture(gl.TEXTURE_2D, depthMap.texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT, depthMap.Width, depthMap.Height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.BindFramebuffer(gl.FRAMEBUFFER, depthMap.fbo)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.TEXTURE_2D, depthMap.texture, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (depthMap *DepthMap) Destroy() {
	gl.DeleteFramebuffers(1, &depthMap.fbo)
}

func (depthMap *DepthMap) Clear() {
	gl.Clear(gl.DEPTH_BUFFER_BIT)
}

func (depthMap *DepthMap) SetActiveShaderProgram(shaderProgram Shader.IShaderProgram) {
	if depthMap.activeShaderProgram != nil {
		depthMap.activeShaderProgram.Unbind()
	}
	depthMap.activeShaderProgram = shaderProgram
	depthMap.activeShaderProgram.Bind()
}

func (depthMap *DepthMap) GetActiveShaderProgram() Shader.IShaderProgram {
	return depthMap.activeShaderProgram
}
