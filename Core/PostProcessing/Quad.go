package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
)

var (
	quadVertices = []Mesh.Vertex{
		{
			Position: GeometryMath.Vector3{-1.0, 1.0, 0.0},
			UV:       GeometryMath.Vector2{0.0, 1.0},
		},
		{
			Position: GeometryMath.Vector3{-1.0, -1.0, 0.0},
			UV:       GeometryMath.Vector2{0.0, 0.0},
		},
		{
			Position: GeometryMath.Vector3{1.0, -1.0, 0.0},
			UV:       GeometryMath.Vector2{1.0, 0.0},
		},
		{
			Position: GeometryMath.Vector3{1.0, 1.0, 0.0},
			UV:       GeometryMath.Vector2{1.0, 1.0},
		},
	}
	quadIndices = []uint32{
		0, 1, 2,
		0, 2, 3,
	}
)

func NewQuad() *Mesh.Mesh {
	return Mesh.NewMesh(quadVertices, quadIndices, nil)
}
