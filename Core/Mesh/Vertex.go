package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Vertex struct {
	Position GeometryMath.Vector3
	UV       GeometryMath.Vector2

	Normal    GeometryMath.Vector3
	Tangent   GeometryMath.Vector3
	BiTangent GeometryMath.Vector3
}

type Vertices []Vertex

func (vertices Vertices) GetPositions() []GeometryMath.Vector3 {
	positions := make([]GeometryMath.Vector3, len(vertices))

	for i, vertex := range vertices {
		positions[i] = vertex.Position
	}

	return positions
}
