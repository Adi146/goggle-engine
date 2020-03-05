package Window

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"testing"
)

func TestSDL(t *testing.T) {
	window := SDLWindow{
		FrameBufferBase: FrameBuffer.FrameBufferBase{
			Width:     100,
			Height:    100,
			DepthTest: false,
			Culling:   false,
			Blend:     false,
		},
		Title: "test window",
		Type: "window",
		Sync: "normal",
	}

	err := window.Init()
	if err != nil {
		t.Errorf("Error while creating sdl window %s", err.Error())
	}

	window.Destroy()
}