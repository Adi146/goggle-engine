package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
	"image"
	"math"
)

type HeightMap struct {
	data       []float32
	NumRows    int
	NumColumns int
	Scale      float32
}

func NewHeightMapFromGrey(img *image.Gray) HeightMap {
	heightMap := HeightMap{
		data:       make([]float32, img.Bounds().Dx()*img.Bounds().Dy()),
		NumRows:    img.Bounds().Dy(),
		NumColumns: img.Bounds().Dx(),
		Scale:      1,
	}

	for z := 0; z < heightMap.NumRows; z++ {
		for x := 0; x < heightMap.NumColumns; x++ {
			heightMap.data[z*heightMap.NumColumns+x] = float32(img.GrayAt(x, z).Y) / math.MaxUint8
		}
	}

	return heightMap
}

func ImportHeightMap(filename string) (HeightMap, error) {
	grey, err := Texture.LoadGray(filename)
	if err != nil {
		return HeightMap{}, err
	}

	return NewHeightMapFromGrey(grey), nil
}

func (heightMap *HeightMap) GetHeight(x int, z int) float32 {
	if x < 0 || x >= heightMap.NumColumns || z < 0 || z >= heightMap.NumRows {
		return 0
	}

	return heightMap.data[z*heightMap.NumColumns+x] * heightMap.Scale
}

func (heightMap *HeightMap) GetNormal(x int, z int) GeometryMath.Vector3 {
	heightL := heightMap.GetHeight(x-1, z)
	heightR := heightMap.GetHeight(x+1, z)
	heightD := heightMap.GetHeight(x, z-1)
	heightU := heightMap.GetHeight(x, z+1)

	return *(&GeometryMath.Vector3{heightL - heightR, 2.0, heightD - heightU}).Normalize()
}

func (heightMap *HeightMap) GetTangent(x int, z int) GeometryMath.Vector3 {
	right := GeometryMath.Vector3{1, 0, 0}
	normal := heightMap.GetNormal(x, z)

	return *right.Cross(&normal)
}

func (heightMap *HeightMap) GetBiTangent(x int, z int) GeometryMath.Vector3 {
	normal := heightMap.GetNormal(x, z)
	tangent := heightMap.GetTangent(x, z)

	return *tangent.Cross(&normal)
}

func (heightMap *HeightMap) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		File  string  `yaml:"file"`
		Scale float32 `yaml:"heightScale"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpHeightMap, err := ImportHeightMap(yamlConfig.File)
	if err != nil {
		return err
	}
	tmpHeightMap.Scale = yamlConfig.Scale

	*heightMap = tmpHeightMap
	return nil
}
