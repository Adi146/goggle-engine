package Node

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const (
	PostProcessingNodeFactoryName   = "Node.PostProcessing"
	PostProcessingShaderFactoryName = "postProcessing"
	OffScreenFramebufferName        = "offScreen"
)

func init() {
	NodeFactory.AddType(PostProcessingNodeFactoryName, reflect.TypeOf((*PostProcessingNode)(nil)).Elem())
	ShaderFactory.AddType(PostProcessingShaderFactoryName, PostProcessing.NewIShaderProgram)
	FrameBufferFactory.AddType(OffScreenFramebufferName, reflect.TypeOf((*FrameBuffer.OffScreenBuffer)(nil)).Elem())
}

type PostProcessingNode struct {
	Scene.INode
	Quad        Model.Mesh
	Shader      ShaderFactory.Config
	FrameBuffer FrameBufferFactory.Config
	Kernel      *PostProcessing.Kernel
}

func (node *PostProcessingNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	if node.Kernel != nil {
		if err := node.Shader.BindObject(node.Kernel); err != nil {
			return err
		}
	}

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *PostProcessingNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if invoker != scene {
		return nil
	}

	node.FrameBuffer.Bind()
	node.FrameBuffer.Clear()
	err := scene.Draw(shader, node, scene)
	node.FrameBuffer.Unbind()
	if err != nil {
		return err
	}

	scene.Clear()

	node.Shader.Bind()
	if shader != nil {
		defer shader.Bind()
	}
	return node.Quad.Draw(node.Shader, node, scene)
}

func (node *PostProcessingNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	yamlConfig := struct {
		Shader      ShaderFactory.Config      `yaml:"shader"`
		FrameBuffer FrameBufferFactory.Config `yaml:"frameBuffer"`
		Kernel      *PostProcessing.Kernel    `yaml:",inline"`
	}{
		Shader:      node.Shader,
		FrameBuffer: node.FrameBuffer,
		Kernel:      node.Kernel,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Shader = yamlConfig.Shader
	node.FrameBuffer = yamlConfig.FrameBuffer
	node.Kernel = yamlConfig.Kernel

	quad, err := PostProcessing.NewQuad()
	if err != nil {
		return err
	}

	node.Quad = *quad

	return nil
}
