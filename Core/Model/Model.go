package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
)

type MeshesWithMaterial struct {
	*Mesh
	*Material
}

type Model struct {
	Meshes      []MeshesWithMaterial
	ModelMatrix *Matrix.Matrix4x4
}
