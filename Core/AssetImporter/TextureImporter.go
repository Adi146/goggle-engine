package AssetImporter

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Model"
	"image"
	"os"
)

func ImportTexture(filename string, textureType Model.TextureType) (*Model.Texture, ImportResult) {
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

	texture, err := Model.NewTextureFromFile(rgba, textureType)
	result.Errors.Push(err)

	return texture, result
}
