package Node

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Skybox"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const SkyboxNodeFactoryName = "Node.Skybox"

func init() {
	Scene.NodeFactory.AddType(SkyboxNodeFactoryName, reflect.TypeOf((*SkyboxNode)(nil)).Elem())
}

type SkyboxNode struct {
	Scene.INode
	Skybox Skybox.Skybox
	Shader Shader.IShaderProgram
}

func (node *SkyboxNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	if scene := node.GetScene(); scene != nil {
		scene.AddOpaqueObject(node)
	}

	return err
}

func (node *SkyboxNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene, camera Camera.ICamera) error {
	if shader == nil {
		node.Shader.Bind()
		defer node.Shader.Unbind()

		shader = node.Shader
	}

	return node.Skybox.Draw(shader, invoker, scene, camera)
}

func (node *SkyboxNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *SkyboxNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Skybox Skybox.Skybox `yaml:"textures"`
		Shader Shader.Ptr    `yaml:"shader"`
	}{
		Skybox: node.Skybox,
		Shader: Shader.Ptr{
			IShaderProgram: node.Shader,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Skybox = yamlConfig.Skybox
	node.Shader = yamlConfig.Shader

	return Scene.UnmarshalChildren(value, node)
}
