package Node

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const ModelNodeFactoryName = "Node.ModelNode"

func init() {
	Scene.NodeFactory.AddType(ModelNodeFactoryName, reflect.TypeOf((*ModelNode)(nil)).Elem())
}

type ModelNode struct {
	Scene.INode
	Model.Model
	IsTransparent bool
	Shader        Shader.IShaderProgram
}

func (node *ModelNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.ModelMatrix = node.GetGlobalTransformation()

	if scene := node.GetScene(); scene != nil {
		if node.IsTransparent {
			scene.AddTransparentObject(node)
		} else {
			scene.AddOpaqueObject(node)
		}
	}

	return err
}

func (node *ModelNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if shader == nil {
		node.Shader.Bind()
		defer node.Shader.Unbind()

		shader = node.Shader
	}

	return node.Model.Draw(shader, nil, nil)
}

func (node *ModelNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	yamlConfig := struct {
		Model         Model.Model `yaml:",inline"`
		IsTransparent bool        `yaml:"isTransparent"`
		Shader        Shader.Ptr  `yaml:"shader"`
	}{
		Model:         node.Model,
		IsTransparent: node.IsTransparent,
		Shader: Shader.Ptr{
			IShaderProgram: node.Shader,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Model = yamlConfig.Model
	node.IsTransparent = yamlConfig.IsTransparent
	node.Shader = yamlConfig.Shader

	return nil
}
