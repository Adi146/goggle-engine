package RenderTarget

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IRenderTarget interface {
	Clear(color *Vector.Vector4)
	Draw(timeDelta float32)
	GetFrameBuffer() Buffer.IFrameBuffer
	SetFrameBuffer(frameBuffer Buffer.IFrameBuffer)
	GetActiveShaderProgram() Shader.IShaderProgram
	SetActiveShaderProgram(shaderProgram Shader.IShaderProgram)
}

func RunRenderLoop(renderer IRenderTarget, window Window.IWindow) {
	for !window.ShouldClose() {
		window.PollEvents()

		timeDelta, _ := window.GetTimeDeltaAndFPS()
		renderer.Draw(timeDelta)
		window.SwapWindow()
	}
}
