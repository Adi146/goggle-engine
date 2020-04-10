package Model

import (
	"fmt"
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
	Material Material.IMaterial
}

type Model struct {
	Meshes       []MeshesWithMaterial
	ModelMatrix  *GeometryMath.Matrix4x4
	NormalMatrix *GeometryMath.Matrix3x3
}

func (model *Model) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(model.ModelMatrix))
	err.Push(shader.BindObject(model.NormalMatrix))
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
		File     string    `yaml:"file"`
		Material yaml.Node `yaml:"material"`
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
		for i := range tmpModel.Meshes {
			if yamlConfig.Material.Kind != 0 {
				if err := yamlConfig.Material.Decode(tmpModel.Meshes[i].Material); err != nil {
					return err
				}
			}
		}

		Log.Warn(&importWarnings, fmt.Sprintf("import warnings: %s", yamlConfig.File))
		*model = *tmpModel
	}

	return importErrors.Err()
}
