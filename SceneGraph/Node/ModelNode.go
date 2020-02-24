package Node

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/AssetImporter"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/Adi146/goggle-engine/Utils/Log"
)

const ModelNodeFactoryName = "Node.ModelNode"

var textureTypeMap = map[string]Texture.TextureType{
	"diffuse":  Texture.DiffuseTexture,
	"specular": Texture.SpecularTexture,
	"emissive": Texture.EmissiveTexture,
	"normals":  Texture.NormalsTexture,
}

func init() {
	YamlFactory.NodeFactory[ModelNodeFactoryName] = reflect.TypeOf((*ModelNodeConfig)(nil)).Elem()
}

type ModelNodeConfig struct {
	Scene.NodeConfig
	File          string              `yaml:"file"`
	Textures      map[string][]string `yaml:"textures"`
	IsTransparent bool                `yaml:"isTransparent"`
	Shader string `yaml:"shader"`
}

func (config *ModelNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &ModelNode{
		INode:  nodeBase,
		Config: config,
	}

	shader, err := ShaderFactory.Get(config.Shader)
	if err != nil {
		return nil, err
	}

	var importErrors Error.ErrorCollection
	var importWarnings Error.ErrorCollection

	model, result := AssetImporter.ImportModel(config.File, shader)
	importErrors.Push(&result.Errors)
	importWarnings.Push(&result.Warnings)
	if result.Success() {
		for textureType, textureFiles := range config.Textures {
			for _, textureFile := range textureFiles {
				texture, result := AssetImporter.ImportTexture(textureFile, textureTypeMap[textureType])
				importErrors.Push(&result.Errors)
				importWarnings.Push(&result.Warnings)
				for _, mesh := range model.Meshes {
					mesh.Textures = append(mesh.Textures, texture)
				}
			}
		}

		Log.Warn(&importWarnings, "import warnings")
		node.Model = model
	}

	return node, importErrors.Err()
}

type ModelNode struct {
	Scene.INode
	*Model.Model

	Config *ModelNodeConfig
}

func (node *ModelNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.ModelMatrix = node.GetGlobalTransformation()

	return err
}

func (node *ModelNode) Draw() error {
	err := node.INode.Draw()

	if scene := node.GetScene(); scene != nil {
		if node.Config.IsTransparent {
			scene.TransparentObjects = append(scene.TransparentObjects, node.Model)
		} else {
			scene.OpaqueObjects = append(scene.OpaqueObjects, node.Model)
		}
	}

	return err
}
