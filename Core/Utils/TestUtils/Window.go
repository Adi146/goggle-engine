package TestUtils

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Window"
	"runtime"
	"testing"
)

func CreateTestWindow(t *testing.T) (*Window.SDLWindow, error) {
	runtime.LockOSThread()

	window, err := Window.NewSDLWindow(FrameBuffer.Viewport{
		PosX:   0,
		PosY:   0,
		Width:  100,
		Height: 100,
	}, "test window", 0)
	if err != nil {
		t.Errorf("Error while creating sdl window %s", err.Error())
	}

	return window, err
}
