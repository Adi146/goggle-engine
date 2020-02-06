package AssetImporter

import (
	"fmt"
	"image"
	"os"

	"github.com/Adi146/goggle-engine/Core/Texture"
)

func ImportTexture(filename string, textureType Texture.TextureType) (*Texture.Texture, ImportResult) {
	var result ImportResult

	file, err := os.Open(filename)
	if err != nil {
		result.Errors.Push(err)
		return nil, result
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		result.Errors.Push(fmt.Errorf("%s: %s", filename, err.Error()))
		return nil, result
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			rgba.Set(x, bounds.Dy()-y, img.At(x, y))
		}
	}

	texture, err := Texture.NewTextureFromFile(rgba, textureType)
	result.Errors.Push(err)

	return texture, result
}
