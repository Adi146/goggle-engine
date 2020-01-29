package RenderTarget

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IRenderTarget interface {
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

		var frameBufferTextures []*Model.Texture
		for i := range frameBuffers {
			frameBuffer := frameBuffers[len(frameBuffers)-1-i]
			renderer.SetFrameBuffer(frameBuffer)

			frameBuffer.Clear()

			for _, frameBufferTexture := range frameBufferTextures {
				err := frameBuffer.GetShaderProgram().BindObject(frameBufferTexture)
				if err != nil {
					fmt.Println(err)
				}
			}
			renderer.Draw()
			frameBufferTextures = frameBuffer.GetTextures()
		}
		window.SwapWindow()
	}
}
