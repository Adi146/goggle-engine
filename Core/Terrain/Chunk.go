package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/go-gl/gl/all-core/gl"
)

const (
	restartIndex = 0xFFFF
)

type Chunk struct {
	Mesh.Mesh
}

func generateChunk(heightMap *HeightMap, offsetX, offsetZ int, chunkRows, chunkColumns int, tileSize float32) Chunk {
	vertexOffsetX := float32(heightMap.NumColumns)/2.0 - float32(offsetX)*tileSize
	vertexOffsetZ := float32(heightMap.NumRows)/2.0 - float32(offsetZ)*tileSize

	vertices := make([]Mesh.Vertex, chunkRows*chunkColumns)
	for z := 0; z < chunkRows; z++ {
		for x := 0; x < chunkColumns; x++ {
			vertices[z*chunkColumns+x] = Mesh.Vertex{
				Position: GeometryMath.Vector3{
					(float32(x) - vertexOffsetX) * tileSize,
					heightMap.GetHeightScaled(x+offsetX, z+offsetZ),
					(float32(z) - vertexOffsetZ) * tileSize,
				},
				Normal:    heightMap.GetNormal(x+offsetX, z+offsetZ),
				UV:        GeometryMath.Vector2{float32(offsetX+x) / float32(heightMap.NumColumns), float32(offsetZ+z) / float32(heightMap.NumRows)},
				Tangent:   heightMap.GetTangent(x+offsetX, z+offsetZ),
				BiTangent: heightMap.GetBiTangent(x+offsetX, z+offsetZ),
			}
		}
	}

	var indices []uint32
	for z := 0; z < chunkRows-1; z++ {
		for x := 0; x < chunkColumns; x++ {
			topLeft := uint32(z*chunkColumns + x)
			bottomLeft := uint32((z+1)*chunkColumns + x)

			indices = append(indices, topLeft, bottomLeft)
		}

		indices = append(indices, restartIndex)
}

	mesh := Mesh.NewMesh(vertices, indices, BoundingVolume.NewBoundingVolumeAABB)
	mesh.EnableFrustumCulling()
	mesh.IndexBuffer.RestartIndex = restartIndex
	mesh.PrimitiveType = gl.TRIANGLE_STRIP

	return Chunk{
		Mesh: *mesh,
	}
}
