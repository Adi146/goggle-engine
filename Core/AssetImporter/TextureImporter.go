package AssetImporter

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Model"
	"image"
	"os"
)

func ImportTexture(filename string) (*Model.Texture, *ImportResult) {
	result := newImportResult()

	file, err := os.Open(filename)
	if err != nil {
		result.addError(err)
		return nil, result
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		result.addError(fmt.Errorf("%s: %s", filename, err.Error()))
		return nil, result
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			rgba.Set(x, bounds.Dy()-y, img.At(x, y))
		}
	}

	texture, err := Model.NewTextureFromFile(rgba)
	if err != nil {
		result.addError(err)
		return nil, result
	}

	result.NumImportedAssets = 1
	return texture, result
}
