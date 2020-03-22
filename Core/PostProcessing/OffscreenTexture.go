package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	OffscreenTexture Texture.Type = "offscreenTexture"
)

func NewOffscreenTexture(width int32, height int32) (*Texture.Texture, error) {
	texture := Texture.Texture{
		Type:   OffscreenTexture,
		Target: gl.TEXTURE_2D,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}
	gl.TexImage2D(texture.Target, 0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	texture.Unbind()

	return &texture, nil
}
