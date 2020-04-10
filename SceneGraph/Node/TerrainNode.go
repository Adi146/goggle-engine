package Node

import (
	"github.com/Adi146/goggle-engine/Core/Model"
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

type TerrainNode ModelNode

func (node *TerrainNode) Tick(timeDelta float32) error {
	return (*ModelNode)(node).Tick(timeDelta)
}

func (node *TerrainNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if shader == nil {
		node.Shader.Bind()
		defer node.Shader.Unbind()

		shader = node.Shader
	}

	return node.Model.Draw(shader, invoker, scene)
}

func (node *TerrainNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
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

	node.Model = (Model.Model)(yamlConfig.Terrain)
	node.Shader = yamlConfig.Shader

	return nil
}
