package RenderTarget

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
	"log"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGLRenderTarget struct {
	Window              Window.IWindow
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
	renderTarget.Window.Destroy()
}

func (renderTarget *OpenGLRenderTarget) GetWindow() Window.IWindow {
	return renderTarget.Window
}

func (renderTarget *OpenGLRenderTarget) Clear(color *Vector.Vector4) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(color[0], color[1], color[2], color[3])

	renderTarget.GetActiveShaderProgram().ResetIndexCounter()
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
	log.Printf("[OpenGL Error] %s", message)
}
