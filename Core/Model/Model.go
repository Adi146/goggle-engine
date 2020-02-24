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
	Shader Shader.IShaderProgram
}

func (model *Model) Draw() error {
	var err Error.ErrorCollection

	model.Shader.Bind()
	err.Push(model.Shader.BindObject(model))
	for _, mesh := range model.Meshes {
		err.Push(model.Shader.BindObject(mesh.Material))
		mesh.Draw()
	}
	model.Shader.Unbind()

	return err.Err()
}

func (model *Model) GetPosition() *Vector.Vector3 {
	return model.ModelMatrix.MulVector(&Vector.Vector3{0, 0, 0})
}
