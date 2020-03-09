package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	Material2 "github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type MeshesWithMaterial struct {
	*Mesh
	*Material2.Material
}

type Model struct {
	Meshes      []MeshesWithMaterial
	ModelMatrix *GeometryMath.Matrix4x4
	Shader      Shader.IShaderProgram
}

func (model *Model) Draw(step *Scene.ProcessingPipelineStep) error {
	var err Error.ErrorCollection

	var shader Shader.IShaderProgram
	if step.EnforcedShader == nil {
		shader = model.Shader
	} else {
		shader = step.EnforcedShader
	}

	shader.Bind()
	err.Push(shader.BindObject(model))
	for _, mesh := range model.Meshes {
		err.Push(shader.BindObject(mesh.Material))
		mesh.Draw()
		mesh.Material.Unbind()
	}
	shader.Unbind()

	return err.Err()
}

func (model *Model) GetPosition() *GeometryMath.Vector3 {
	return model.ModelMatrix.MulVector(&GeometryMath.Vector3{0, 0, 0})
}
