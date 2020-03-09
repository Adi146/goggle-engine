package TestUtils

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Window"
	"testing"
)

func CreateTestWindow(t *testing.T) (Window.SDLWindow, error) {
	window := Window.SDLWindow{
		FrameBufferBase: FrameBuffer.FrameBufferBase{
			Width:     100,
			Height:    100,
			DepthTest: false,
			Culling: FrameBuffer.CullFunction{
				Enabled:  false,
				Function: 0,
			},
			Blend: false,
		},
		Title: "test window",
		Type:  "window",
		Sync:  "normal",
	}

	err := window.Init()
	if err != nil {
		t.Errorf("Error while creating sdl window %s", err.Error())
	}

	return window, err
}
