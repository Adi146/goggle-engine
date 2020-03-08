package Texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type TextureType string
type TextureTarget uint32

type Texture struct {
	TextureID     uint32
	TextureType   TextureType
	TextureTarget TextureTarget
}

func (tex *Texture) Bind(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(uint32(tex.TextureTarget), tex.TextureID)
}

func (tex *Texture) Unbind() {
	gl.BindTexture(uint32(tex.TextureTarget), 0)
}
