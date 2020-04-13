package Terrain

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type ITerrain interface {
	GetHeightAt(terrainPos GeometryMath.Vector3) float32
}
