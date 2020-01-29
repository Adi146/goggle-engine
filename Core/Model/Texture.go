package Model

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	_ "image/png"
)

type TextureType string

const (
	DiffuseTexture  TextureType = "diffuse"
	SpecularTexture TextureType = "specular"
	EmissiveTexture TextureType = "emissive"
	NormalsTexture  TextureType = "normals"
)

type Texture struct {
	TextureID   uint32
	TextureType TextureType
}

func NewTextureFromFile(img *image.RGBA, textureType TextureType) (*Texture, error) {
	texture := Texture{
		TextureType: textureType,
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

func (tex *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tex.TextureID)
}

func (tex *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
