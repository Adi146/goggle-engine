package AssetImporter

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/Adi146/goggle-engine/Core/Texture"
)

type CubeMapImportHelper struct {
	Top    string `yaml:"top"`
	Bottom string `yaml:"bottom"`
	Left   string `yaml:"left"`
	Right  string `yaml:"right"`
	Front  string `yaml:"front"`
	Back   string `yaml:"back"`
}

func (helper *CubeMapImportHelper) getArray() []string {
	return []string{
		helper.Right,
		helper.Left,
		helper.Top,
		helper.Bottom,
		helper.Front,
		helper.Back,
	}
}

func ImportTexture(filename string, textureType Texture.Type) (*Texture.Texture, ImportResult) {
	var result ImportResult

	rgba, err := loadRGBA(filename)
	if err != nil {
		result.Errors.Push(err)
		return nil, result
	}

	rgba = flipRGBA(rgba)
	texture, err := Texture.NewTextureFromFile(rgba, textureType)
	result.Errors.Push(err)

	return texture, result
}

func ImportCubeMap(files CubeMapImportHelper, textureType Texture.Type) (*Texture.Texture, ImportResult) {
	var result ImportResult
	rgbas := make([]*image.RGBA, len(files.getArray()))

	for i, filename := range files.getArray() {
		rgba, err := loadRGBA(filename)
		if err != nil {
			result.Errors.Push(err)
		} else {
			rgbas[i] = rgba
		}
	}

	cubeMap, err := Texture.NewCubeMapFromFile(rgbas, textureType)
	result.Errors.Push(err)

	return cubeMap, result
}

func loadRGBA(filename string) (*image.RGBA, error) {
	img, err := loadImage(filename)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)

	return rgba, nil
}

func flipRGBA(rgba *image.RGBA) *image.RGBA {
	bounds := rgba.Bounds()
	newRgba := image.NewRGBA(bounds)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			newRgba.SetRGBA(x, bounds.Dy()-y, rgba.RGBAAt(x, y))
		}
	}

	return newRgba
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", filename, err.Error())
	}

	return img, nil
}
