package VertexBuffer

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Vertex struct {
	Position GeometryMath.Vector3
	UV      GeometryMath.Vector2

	Normal   GeometryMath.Vector3
	Tangent GeometryMath.Vector3
	BiTangent GeometryMath.Vector3
}
