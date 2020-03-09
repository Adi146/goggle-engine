package Texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Unit struct {
	ID       uint32
	Texture  ITexture
}

func (unit *Unit) Activate() {
	gl.ActiveTexture(gl.TEXTURE0 + unit.ID)
}