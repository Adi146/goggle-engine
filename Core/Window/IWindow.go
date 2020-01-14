package Window

type IWindow interface {
	Init() error
	Destroy()

	SwapWindow()
	PollEvents()
	ShouldClose() bool

	GetTimeDeltaAndFPS() (float32, uint32)
	GetSize() (int32, int32)

	GetKeyboardInput() IKeyboardInput
	GetMouseInput() IMouseInput

	EnableVSync()
	EnableAdaptiveSync()
}
