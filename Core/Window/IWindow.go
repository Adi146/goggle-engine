package Window

import "github.com/Adi146/goggle-engine/Core/Buffer"

type IWindow interface {
	Buffer.IFrameBuffer

	Init() error

	SwapWindow()
	PollEvents()
	ShouldClose() bool

	GetTimeDeltaAndFPS() (float32, uint32)

	GetKeyboardInput() IKeyboardInput
	GetMouseInput() IMouseInput

	EnableVSync()
	EnableAdaptiveSync()
}
