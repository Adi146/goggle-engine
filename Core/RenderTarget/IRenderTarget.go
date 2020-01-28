package RenderTarget

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IRenderTarget interface {
	Clear(color *Vector.Vector4)
	Tick(timeDelta float32)
	Draw()
	GetFrameBuffer() FrameBuffer.IFrameBuffer
	SetFrameBuffer(frameBuffer FrameBuffer.IFrameBuffer)
}

func RunRenderLoop(renderer IRenderTarget, frameBuffers []FrameBuffer.IFrameBuffer) {
	window := frameBuffers[0].(Window.IWindow)

	for !window.ShouldClose() {
		window.PollEvents()

		timeDelta, _ := window.GetTimeDeltaAndFPS()
		renderer.Tick(timeDelta)
		for _, frameBuffer := range frameBuffers {
			renderer.SetFrameBuffer(frameBuffer)
			renderer.Draw()
		}
		window.SwapWindow()
	}
}
