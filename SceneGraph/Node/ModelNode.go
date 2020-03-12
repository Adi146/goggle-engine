package Node

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const ModelNodeFactoryName = "Node.ModelNode"

var textureTypeMap = map[string]Texture.Type{
	"diffuse":  Texture.DiffuseTexture,
	"specular": Texture.SpecularTexture,
	"emissive": Texture.EmissiveTexture,
	"normals":  Texture.NormalsTexture,
}

func init() {
	NodeFactory.AddType(ModelNodeFactoryName, reflect.TypeOf((*ModelNodeConfig)(nil)).Elem())
}

type ModelNodeConfig struct {
	Scene.NodeConfig
	Model         Model.Model          `yaml:",inline"`
	IsTransparent bool                 `yaml:"isTransparent"`
	Shader        ShaderFactory.Config `yaml:"shader"`
}

func (config *ModelNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &ModelNode{
		INode:  nodeBase,
		Model:  &config.Model,
		Config: config,
	}

	return node, nil
}

type ModelNode struct {
	Scene.INode
	*Model.Model

	Config *ModelNodeConfig
}

func (node *ModelNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.ModelMatrix = node.GetGlobalTransformation()

	if scene := node.GetScene(); scene != nil {
		if node.Config.IsTransparent {
			scene.AddTransparentObject(node)
		} else {
			scene.AddOpaqueObject(node)
		}
	}

	return err
}

func (node *ModelNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if shader == nil {
		node.Config.Shader.Bind()
		defer node.Config.Shader.Unbind()

		shader = node.Config.Shader
	}

	return node.Model.Draw(shader, nil, nil)
}
