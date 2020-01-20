package Node

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/AssetImporter"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
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

func (node *ModelNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	if node.Model == nil {
		model, result := AssetImporter.ImportModel(node.File)
		if !result.Success() {
			return fmt.Errorf("errors while importing model")
		}

		for _, diffuseTextureFile := range node.Textures.Diffuse {
			texture, result := AssetImporter.ImportTexture(diffuseTextureFile)
			if !result.Success() {
				return fmt.Errorf("errors while importing texture %s", diffuseTextureFile)
			}
			for _, mesh := range model.Meshes {
				mesh.DiffuseTextures = append(mesh.DiffuseTextures, texture)
			}
		}

		for _, normalsTextureFile := range node.Textures.Normals {
			texture, result := AssetImporter.ImportTexture(normalsTextureFile)
			if !result.Success() {
				return fmt.Errorf("errors while importing texture %s", normalsTextureFile)
			}
			for _, mesh := range model.Meshes {
				mesh.NormalTextures = append(mesh.NormalTextures, texture)
			}
		}

		node.Model = model
	}

	return nil
}

func (node *ModelNode) Tick(timeDelta float32) {
	node.ModelMatrix = node.GetGlobalTransformation()
}

func (node *ModelNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindObject(node.Model)
		for _, mesh := range node.Meshes {
			scene.GetActiveShaderProgram().BindObject(mesh.Material)
			mesh.Draw()
		}
	}
}
