package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type MeshesWithMaterial struct {
	*Mesh
	*Material.Material
}

type Model struct {
	Meshes      []MeshesWithMaterial
	ModelMatrix *GeometryMath.Matrix4x4
}

func (model *Model) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(model))
	for _, mesh := range model.Meshes {
		err.Push(shader.BindObject(mesh.Material))
		mesh.Draw(shader, nil, nil)
		mesh.Material.Unbind()
	}

	return err.Err()
}

func (model *Model) GetPosition() *GeometryMath.Vector3 {
	return model.ModelMatrix.MulVector(&GeometryMath.Vector3{0, 0, 0})
}
