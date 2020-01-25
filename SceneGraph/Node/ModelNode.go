package Node

import (
	"github.com/Adi146/goggle-engine/Core/AssetImporter"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"reflect"
)

type TextureConfiguration struct {
	Diffuse []string `yaml:"diffuse"`
	Normals []string `yaml:"normals"`
}

type ModelNode struct {
	Scene.IChildNode
	*Model.Model

	File     string               `yaml:"file"`
	Textures TextureConfiguration `yaml:"textures"`
}

func init() {
	Factory.NodeFactory["Node.ModelNode"] = reflect.TypeOf((*ModelNode)(nil)).Elem()
}

func (node *ModelNode) Init(nodeID string) error {
	var err Error.ErrorCollection

	if node.IChildNode == nil {
		node.IChildNode = &Scene.ChildNodeBase{}
		if err := node.IChildNode.Init(nodeID); err != nil {
			return err
		}
	}

	if node.Model == nil {
		var importWarnings Error.ErrorCollection

		model, result := AssetImporter.ImportModel(node.File)
		err.Push(&result.Errors)
		importWarnings.Push(&result.Warnings)
		if result.Success() {
			for _, diffuseTextureFile := range node.Textures.Diffuse {
				texture, result := AssetImporter.ImportTexture(diffuseTextureFile, Model.DiffuseTexture)
				err.Push(&result.Errors)
				importWarnings.Push(&result.Warnings)
				for _, mesh := range model.Meshes {
					mesh.Textures = append(mesh.Textures, texture)
				}
			}

			for _, normalsTextureFile := range node.Textures.Normals {
				texture, result := AssetImporter.ImportTexture(normalsTextureFile, Model.NormalsTexture)
				err.Push(&result.Errors)
				importWarnings.Push(&result.Warnings)
				for _, mesh := range model.Meshes {
					mesh.Textures = append(mesh.Textures, texture)
				}
			}

			Log.Warn(Error.NewErrorWithFields(&importWarnings, node.GetLogFields()), "import warnings")
			node.Model = model
		}
	}

	return err.Err()
}

func (node *ModelNode) Tick(timeDelta float32) error {
	node.ModelMatrix = node.GetGlobalTransformation()
	return nil
}

func (node *ModelNode) Draw() error {
	var err Error.ErrorCollection

	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		err.Push(scene.GetActiveShaderProgram().BindObject(node.Model))
		for _, mesh := range node.Meshes {
			err.Push(scene.GetActiveShaderProgram().BindObject(mesh.Material))
			mesh.Draw()
		}
	}

	return err.Err()
}
