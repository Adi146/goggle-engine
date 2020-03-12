package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IScene interface {
	IDrawable

	Init() error
	Tick(timeDelta float32)

	GetKeyboardInput() Window.IKeyboardInput
	SetKeyboardInput(input Window.IKeyboardInput)
	GetMouseInput() Window.IMouseInput
	SetMouseInput(input Window.IMouseInput)
}
