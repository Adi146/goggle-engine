package Terrain

import (
	"gopkg.in/yaml.v3"
	"image"
)

type PlacementMap HeightMap

func NewPlacementMapFromGrey(img *image.Gray) PlacementMap {
	return (PlacementMap)(NewHeightMapFromGrey(img))
}

func ImportPlacementMap(filename string) (PlacementMap, error) {
	heightMap, err := ImportHeightMap(filename)
	return (PlacementMap)(heightMap), err
}

func (placementMap *PlacementMap) GetProbability(x int, z int) float32 {
	return (*HeightMap)(placementMap).GetHeight(x, z)
}

func (placementMap *PlacementMap) GetProbabilityAtArea(minX, minZ, maxX, maxZ int) float32 {
	var sumProbability float32
	var numProbabilities float32

	for z := minZ; z < maxZ; z++ {
		for x := minX; x < maxX; x++ {
			probability := placementMap.GetProbability(x, z)
			if probability == 0 {
				return 0.0
			}

			sumProbability += probability
			numProbabilities++
		}
	}

	return sumProbability / numProbabilities
}

func (placementMap *PlacementMap) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig string
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpPlacementMap, err := ImportPlacementMap(yamlConfig)
	if err != nil {
		return err
	}

	*placementMap = tmpPlacementMap
	return nil
}
