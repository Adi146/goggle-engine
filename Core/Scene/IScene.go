package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type IScene interface {
	IDrawable

	Start() error
	Tick(timeDelta float32)

	Clear()
	AddPreRenderObject(obj IDrawable)
	AddOpaqueObject(obj IDrawable)
	AddTransparentObject(obj ITransparentDrawable)

	GetWindow() Window.IWindow
	GetCamera() Camera.ICamera
}
