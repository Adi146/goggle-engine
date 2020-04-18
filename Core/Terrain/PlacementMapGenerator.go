package Terrain

type PlacementMapGenerator struct {
	Ranges []Range `yaml:"ranges"`
}

//func (generator *PlacementMapGenerator) GeneratePlacementMap(heightMap *HeightMap) *PlacementMap {
//	img := image.NewGray(image.Rect(0, 0, heightMap.NumColumns, heightMap.NumRows))
//
//	for z := 0; z < heightMap.NumRows; z ++ {
//		for x := 0 ; x < heightMap.NumColumns; x ++ {
//			currentHeight := heightMap.GetHeight(x, z)
//
//
//		}
//	}
//}
