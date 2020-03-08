package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Scene"
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
	}
	shader.Unbind()

	return err.Err()
}

func (model *Model) GetPosition() *Vector.Vector3 {
	return model.ModelMatrix.MulVector(&Vector.Vector3{0, 0, 0})
}
