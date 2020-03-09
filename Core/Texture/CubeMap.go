package Texture

import (
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	SkyboxTexture Type = "skybox"
)

func NewCubeMapFromFile(images []*image.RGBA, textureType Type) (*Texture, error) {
	texture := Texture{
		Target: gl.TEXTURE_CUBE_MAP,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}

	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	for i, img := range images {
		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i), 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	}
	texture.Unbind()

	return &texture, nil
}
