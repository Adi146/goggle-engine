package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IScene interface {
	Init() error
	Tick(timeDelta float32)
	Draw(step *ProcessingPipelineStep)

	GetKeyboardInput() Window.IKeyboardInput
	SetKeyboardInput(input Window.IKeyboardInput)
	GetMouseInput() Window.IMouseInput
	SetMouseInput(input Window.IMouseInput)
}
