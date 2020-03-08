package Texture

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	DiffuseTexture  TextureType = "diffuse"
	SpecularTexture TextureType = "specular"
	EmissiveTexture TextureType = "emissive"
	NormalsTexture  TextureType = "normals"

	Texture2D TextureTarget = gl.TEXTURE_2D
)

func NewTextureFromFile(img *image.RGBA, textureType TextureType) (*Texture, error) {
	texture := Texture{
		TextureType:   textureType,
		TextureTarget: Texture2D,
	}

	gl.GenTextures(1, &texture.TextureID)
	gl.BindTexture(gl.TEXTURE_2D, texture.TextureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return &texture, nil
}
