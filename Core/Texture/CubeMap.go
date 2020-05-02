package Texture

import (
	"github.com/Adi146/goggle-engine/Core/Utils"
	"image"

	"github.com/go-gl/gl/v4.3-core/gl"
)

const (
	SkyboxTexture Type = "skybox"
)

type CubeMap struct {
	Texture
}

func NewCubeMapFromRGBAs(images []*image.RGBA, textureType Type) (*CubeMap, error) {
	texture := Texture{
		Target: gl.TEXTURE_CUBE_MAP,
	}

	gl.GenTextures(1, &texture.ID)
	if err := texture.Bind(); err != nil {
		return nil, err
	}

	gl.TexParameteri(texture.Target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(texture.Target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	texture.SetWrapMode(ClampToEdge)

	for i, img := range images {
		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i), 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, Utils.GlPtr(img.Pix))
	}
	texture.Unbind()

	return &CubeMap{
		Texture: texture,
	}, nil
}

func ImportCubeMap(files []string, textureType Type) (*CubeMap, error) {
	rgbas := make([]*image.RGBA, len(files))

	for i, filename := range files {
		rgba, err := loadRGBA(filename)
		if err != nil {
			return nil, err
		}
		rgbas[i] = rgba
	}

	cubeMap, err := NewCubeMapFromRGBAs(rgbas, textureType)
	if err != nil {
		return nil, err
	}

	return cubeMap, nil
}
