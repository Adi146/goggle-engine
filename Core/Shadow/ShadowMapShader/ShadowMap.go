package ShadowMapShader

import (
	"github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	ShadowMap Texture.Type = "shadowMap"
)

func NewShadowMap(width int32, height int32) (*Texture.Texture, error) {
	texture := Texture.Texture{
		Type:   ShadowMap,
		Target: gl.TEXTURE_2D,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}
	gl.TexImage2D(texture.Target, 0, gl.DEPTH_COMPONENT24, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	texture.Unbind()

	return &texture, nil
}

func NewShadowCubeMap(width int32, height int32) (*Texture.Texture, error) {
	texture := Texture.Texture{
		Type:   ShadowMap,
		Target: gl.TEXTURE_CUBE_MAP,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}

	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(texture.Target, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	for i := uint32(0); i < 6; i++ {
		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+i, 0, gl.DEPTH_COMPONENT24, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	}
	texture.Unbind()

	return &texture, nil
}
