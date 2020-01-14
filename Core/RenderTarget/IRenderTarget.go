package RenderTarget

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IRenderTarget interface {
	Clear(color *Vector.Vector4)
	Draw(timeDelta float32)
	GetWindow() Window.IWindow
	GetActiveShaderProgram() Shader.IShaderProgram
	SetActiveShaderProgram(shaderProgram Shader.IShaderProgram)
}

func RunRenderLoop(renderer IRenderTarget) {
	for !renderer.GetWindow().ShouldClose() {
		renderer.GetWindow().PollEvents()

		timeDelta, _ := renderer.GetWindow().GetTimeDeltaAndFPS()
		renderer.Draw(timeDelta)
		renderer.GetWindow().SwapWindow()
	}
}
