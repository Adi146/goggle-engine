package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"image"
	"image/color"
)

type BlendMapGenerator struct {
	R Range `yaml:"r"`
	G Range `yaml:"g"`
	B Range `yaml:"b"`
}

type Range struct {
	Min float32 `yaml:"min"`
	Max float32 `yaml:"max"`
}

func (generator *BlendMapGenerator) GenerateBlendMap(heightMap *HeightMap) (*Texture.Texture2D, error) {
	img := image.NewRGBA(image.Rect(0, 0, heightMap.NumColumns, heightMap.NumRows))

	for z := 0; z < heightMap.NumRows; z++ {
		for x := 0; x < heightMap.NumColumns; x++ {
			currentHeight := heightMap.GetHeightScaled(x, z)

			pixelColor := color.RGBA{
				R: uint8(generator.R.GetFactor(currentHeight) * 255),
				G: uint8(generator.G.GetFactor(currentHeight) * 255),
				B: uint8(generator.B.GetFactor(currentHeight) * 255),
				A: 255,
			}

			img.Set(x, z, pixelColor)
		}
	}

	return Texture.NewTextureFromRGBA(img, Material.BlendMap)
}

func (heightRange *Range) GetFactor(height float32) float32 {
	if height < heightRange.Min {
		return height / heightRange.Min
	}

	if height > heightRange.Max {
		return heightRange.Max / height
	}

	return 1
}
