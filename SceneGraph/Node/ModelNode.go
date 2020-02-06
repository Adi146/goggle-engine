package Node

import (
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
	Scene.ChildNodeBaseConfig
	File     string              `yaml:"file"`
	Textures map[string][]string `yaml:"textures"`
}

func (config ModelNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsChildNode()
}

func (config ModelNodeConfig) CreateAsChildNode() (Scene.IChildNode, error) {
	nodeBase, err := config.ChildNodeBaseConfig.CreateAsChildNode()
	if err != nil {
		return nil, err
	}

	node := &ModelNode{
		ModelNodeConfig: &config,
		IChildNode:      nodeBase,
	}

	var importErrors Error.ErrorCollection
	var importWarnings Error.ErrorCollection

	model, result := AssetImporter.ImportModel(config.File)
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

		Log.Warn(Error.NewErrorWithFields(&importWarnings, node.GetLogFields()), "import warnings")
		node.Model = model
	}

	return node, importErrors.Err()
}

type ModelNode struct {
	*ModelNodeConfig
	Scene.IChildNode
	*Model.Model
}

func (node *ModelNode) Tick(timeDelta float32) error {
	err := node.IChildNode.Tick(timeDelta)

	node.ModelMatrix = node.GetGlobalTransformation()

	return err
}

func (node *ModelNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.OpaqueObjects = append(scene.OpaqueObjects, node.Model)
	}

	return nil
}
