package Window

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
)

type IWindow interface {
	FrameBuffer.IFrameBuffer

	SwapWindow()
	PollEvents()
	ShouldClose() bool

	GetTimeDeltaAndFPS() (float32, uint32)

	GetKeyboardInput() IKeyboardInput
	GetMouseInput() IMouseInput

	EnableVSync()
	EnableAdaptiveSync()
}
