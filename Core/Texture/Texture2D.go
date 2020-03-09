package Texture

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	DiffuseTexture  Type = "diffuse"
	SpecularTexture Type = "specular"
	EmissiveTexture Type = "emissive"
	NormalsTexture  Type = "normals"
)

func NewTextureFromFile(img *image.RGBA, textureType Type) (*Texture, error) {
	texture := Texture{
		Type:   textureType,
		Target: gl.TEXTURE_2D,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}

	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(texture.Target, 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	texture.Unbind()

	return &texture, nil
}
