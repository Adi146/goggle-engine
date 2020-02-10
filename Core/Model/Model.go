package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type MeshesWithMaterial struct {
	*Mesh
	*Material
}

type Model struct {
	Meshes      []MeshesWithMaterial
	ModelMatrix *Matrix.Matrix4x4
}

func (model *Model) Draw(shader Shader.IShaderProgram) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(model))
	for _, mesh := range model.Meshes {
		err.Push(shader.BindObject(mesh.Material))
		mesh.Draw()
	}

	return err.Err()
}

func (model *Model) GetPosition() *Vector.Vector3 {
	return model.ModelMatrix.MulVector(&Vector.Vector3{0, 0, 0})
}
