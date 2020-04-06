package Texture

import (
	"fmt"
	"image"
	"image/draw"
	"os"
)

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

func LoadGray(filename string) (*image.Gray, error) {
	img, err := loadImage(filename)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	draw.Draw(gray, gray.Bounds(), img, bounds.Min, draw.Src)

	return gray, nil
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
