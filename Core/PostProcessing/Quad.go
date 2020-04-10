package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/VertexBuffer"
)

var (
	quadVertices = []VertexBuffer.Vertex{
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

func NewQuad() (*Model.Mesh, error) {
	return Model.NewMesh(quadVertices, quadIndices)
}
