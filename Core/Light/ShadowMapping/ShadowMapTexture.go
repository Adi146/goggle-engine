package ShadowMapping

import (
	core "github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/go-gl/gl/v4.3-core/gl"
)

const (
	ShadowMapDirectionalLight core.Type = "shadowMapDirectionalLight"
	ShadowMapPointLight       core.Type = "shadowMapPointLight"
	ShadowMapSpotLight        core.Type = "shadowMapSpotLight"
)

func NewShadowMapTexture(width int32, height int32, textureType core.Type) (*core.Texture, error) {
	texture := core.Texture{
		Type:   textureType,
		Target: gl.TEXTURE_2D,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}
	gl.TexImage2D(texture.Target, 0, gl.DEPTH_COMPONENT24, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	texture.SetWrapMode(core.ClampToEdge)

	texture.Unbind()

	return &texture, nil
}

func NewShadowCubeMapTexture(width int32, height int32, textureType core.Type) (*core.Texture, error) {
	texture := core.Texture{
		Type:   textureType,
		Target: gl.TEXTURE_CUBE_MAP,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}
	for i := uint32(0); i < 6; i++ {
		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+i, 0, gl.DEPTH_COMPONENT24, width, height, 0, gl.DEPTH_COMPONENT, gl.FLOAT, nil)
	}
	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	texture.SetWrapMode(core.ClampToEdge)
	texture.Unbind()

	return &texture, nil
}
