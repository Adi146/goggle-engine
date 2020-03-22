package FrameBuffer

import "github.com/go-gl/gl/v4.1-core/gl"

type Viewport struct {
	PosX   int32 `yaml:"x-position"`
	PosY   int32 `yaml:"y-position"`
	Width  int32 `yaml:"width"`
	Height int32 `yaml:"height"`
}

func (viewport *Viewport) Bind() {
	gl.Viewport(viewport.PosX, viewport.PosY, viewport.Width, viewport.Height)
}

func GetCurrentViewport() Viewport {
	var viewport [4]int32

	gl.GetIntegerv(gl.VIEWPORT, &viewport[0])

	return Viewport{
		PosX:   viewport[0],
		PosY:   viewport[1],
		Width:  viewport[2],
		Height: viewport[3],
	}
}
