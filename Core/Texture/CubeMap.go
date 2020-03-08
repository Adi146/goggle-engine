package Texture

import (
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	SkyboxTexture TextureType = "skybox"

	CubeMap TextureTarget = gl.TEXTURE_CUBE_MAP
)

func NewCubeMapFromFile(images []*image.RGBA, textureType TextureType) (*Texture, error) {
	texture := Texture{
		TextureTarget: CubeMap,
	}

	gl.GenTextures(1, &texture.TextureID)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, texture.TextureID)

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	for i, img := range images {
		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i), 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	}
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, 0)

	return &texture, nil
}
