package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
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

	err.Push(shader.BindObject(model.ModelMatrix))
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

func (model *Model) UnmarshalYAML(value *yaml.Node) error {
	var importErrors Error.ErrorCollection
	var importWarnings Error.ErrorCollection

	var yamlConfig struct {
		File     string            `yaml:"file"`
		Material Material.Material `yaml:",inline"`
	}
	var err error
	if value.Kind == yaml.ScalarNode {
		err = value.Decode(&yamlConfig.File)
	} else {
		err = value.Decode(&yamlConfig)
	}
	if err != nil {
		return err
	}

	tmpModel, result := ImportModel(yamlConfig.File)
	importErrors.Push(&result.Errors)
	importWarnings.Push(&result.Warnings)
	if result.Success() {
		for _, mesh := range tmpModel.Meshes {
			mesh.Textures = append(mesh.Textures, yamlConfig.Material.Textures...)
		}

		Log.Warn(&importWarnings, "import warnings")
		*model = *tmpModel
	}

	return importErrors.Err()
}
