package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IScene interface {
	Init() error
	Tick(timeDelta float32)
	Draw()

	GetKeyboardInput() Window.IKeyboardInput
	SetKeyboardInput(input Window.IKeyboardInput)
	GetMouseInput() Window.IMouseInput
	SetMouseInput(input Window.IMouseInput)

	AddResult(texture *Texture.Texture)
}
