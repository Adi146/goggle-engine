package RenderTarget

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGLRenderTarget struct {
	frameBuffer FrameBuffer.IFrameBuffer

	Culling      bool `yaml:"culling"`
	DepthTest    bool `yaml:"depthTest"`
	DebugLogging bool `yaml:"debugLogging"`
}

func (renderTarget *OpenGLRenderTarget) Init() error {
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
	renderTarget.frameBuffer.Destroy()
}

func (renderTarget *OpenGLRenderTarget) GetFrameBuffer() FrameBuffer.IFrameBuffer {
	return renderTarget.frameBuffer
}

func (renderTarget *OpenGLRenderTarget) SetFrameBuffer(frameBuffer FrameBuffer.IFrameBuffer) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, frameBuffer.GetFBO())
	width, height := frameBuffer.GetSize()
	gl.Viewport(0, 0, width, height)

	frameBuffer.GetShaderProgram().Bind()

	renderTarget.frameBuffer = frameBuffer
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
