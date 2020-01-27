package RenderTarget

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGLRenderTarget struct {
	frameBuffer         Buffer.IFrameBuffer
	activeShaderProgram Shader.IShaderProgram

	Culling      bool `yaml:"culling"`
	DepthTest    bool `yaml:"depthTest"`
	DebugLogging bool `yaml:"debugLogging"`
}

func (renderTarget *OpenGLRenderTarget) Init() error {
	if err := gl.Init(); err != nil {
		return err
	}
	if renderTarget.Culling {
		EnableCulling()
	}
	if renderTarget.DepthTest {
		EnableDepthTest()
	}
	if renderTarget.DebugLogging {
		EnableDebugLogging()
	}

	return nil
}

func (renderTarget *OpenGLRenderTarget) Destroy() {
	renderTarget.activeShaderProgram.Destroy()
	renderTarget.frameBuffer.Destroy()
}

func (renderTarget *OpenGLRenderTarget) GetFrameBuffer() Buffer.IFrameBuffer {
	return renderTarget.frameBuffer
}

func (renderTarget *OpenGLRenderTarget) SetFrameBuffer(frameBuffer Buffer.IFrameBuffer) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, frameBuffer.GetFBO())
	width, height := frameBuffer.GetSize()
	gl.Viewport(0, 0, width, height)

	renderTarget.frameBuffer = frameBuffer
}

func (renderTarget *OpenGLRenderTarget) Clear(color *Vector.Vector4) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(color[0], color[1], color[2], color[3])
}

func (renderTarget *OpenGLRenderTarget) SetActiveShaderProgram(shaderProgram Shader.IShaderProgram) {
	if renderTarget.activeShaderProgram != nil {
		renderTarget.activeShaderProgram.Unbind()
	}
	renderTarget.activeShaderProgram = shaderProgram
	renderTarget.activeShaderProgram.Bind()
}

func (renderTarget *OpenGLRenderTarget) GetActiveShaderProgram() Shader.IShaderProgram {
	return renderTarget.activeShaderProgram
}

func EnableDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
}

func EnableCulling() {
	gl.Enable(gl.CULL_FACE)
}

func EnableDebugLogging() {
	gl.Enable(gl.DEBUG_OUTPUT)
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	gl.DebugMessageCallback(openGLDebugCallback, nil)
}

func openGLDebugCallback(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	Log.Error(fmt.Errorf("[OpenGL Error] %s", message), "OpenGL Error")
}
