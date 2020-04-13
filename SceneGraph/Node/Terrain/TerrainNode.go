package Terrain

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Terrain"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const TerrainNodeFactoryName = "Node.Terrain"

func init() {
	Scene.NodeFactory.AddType(TerrainNodeFactoryName, reflect.TypeOf((*TerrainNode)(nil)).Elem())
}

type TerrainNode struct {
	Scene.INode
	Terrain.Terrain
	Shader Shader.IShaderProgram
}

func (node *TerrainNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetModelMatrix(node.GetGlobalTransformation())

	if scene := node.GetScene(); scene != nil {
		scene.AddOpaqueObject(node)
	}

	return err
}

func (node *TerrainNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if shader == nil {
		node.Shader.Bind()
		defer node.Shader.Unbind()

		shader = node.Shader
	}

	return node.Model.Draw(shader, invoker, scene)
}

func (node *TerrainNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *TerrainNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	yamlConfig := struct {
		Shader  Shader.Ptr      `yaml:"shader"`
		Terrain Terrain.Terrain `yaml:",inline"`
	}{
		Shader: Shader.Ptr{
			IShaderProgram: node.Shader,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Terrain = yamlConfig.Terrain
	node.Shader = yamlConfig.Shader

	return Scene.UnmarshalChildren(value, node, AnchorNodeFactoryName)
}
