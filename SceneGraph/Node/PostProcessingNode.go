package Node

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const (
	PostProcessingNodeFactoryName = "Node.PostProcessing"
)

func init() {
	Scene.NodeFactory.AddType(PostProcessingNodeFactoryName, reflect.TypeOf((*PostProcessingNode)(nil)).Elem())
}

type PostProcessingNode struct {
	Scene.INode
	Quad        Mesh.Mesh
	Shader      Shader.IShaderProgram
	FrameBuffer FrameBuffer.FrameBuffer
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

	oldFrameBuffer := FrameBuffer.GetCurrentFrameBuffer()
	node.FrameBuffer.Bind()
	node.FrameBuffer.Clear()

	err := scene.Draw(shader, node, scene)
	oldFrameBuffer.Bind()
	if err != nil {
		return err
	}

	defer Function.GetCurrentCullFunction().Set()
	defer Function.GetCurrentDepthFunction().Set()
	defer Function.GetCurrentBlendFunction().Set()

	Function.Back.Set()
	Function.DisabledDepth.Set()
	Function.DisabledBlend.Set()

	node.Shader.Bind()
	if shader != nil {
		defer shader.Bind()
	}

	scene.Clear()
	return node.Quad.Draw(node.Shader, node, scene)
}

func (node *PostProcessingNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *PostProcessingNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Shader      Shader.Ptr              `yaml:"shader"`
		FrameBuffer FrameBuffer.FrameBuffer `yaml:"frameBuffer"`
		Kernel      *PostProcessing.Kernel  `yaml:",inline"`
	}{
		Shader: Shader.Ptr{
			IShaderProgram: node.Shader,
		},
		FrameBuffer: node.FrameBuffer,
		Kernel:      node.Kernel,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	texture, err := PostProcessing.NewOffscreenTexture(yamlConfig.FrameBuffer.Viewport.Width, yamlConfig.FrameBuffer.Viewport.Height)
	if err != nil {
		return err
	}
	yamlConfig.FrameBuffer.AddColorAttachment(texture, 0)
	yamlConfig.FrameBuffer.AddDepthStencilAttachment(FrameBuffer.NewDepth24Stencil8Rbo(yamlConfig.FrameBuffer.Viewport.Width, yamlConfig.FrameBuffer.Viewport.Height))
	if err := yamlConfig.FrameBuffer.Finish(); err != nil {
		return err
	}

	if err := yamlConfig.Shader.BindObject(texture); err != nil {
		return err
	}

	node.Shader = yamlConfig.Shader
	node.FrameBuffer = yamlConfig.FrameBuffer
	node.Kernel = yamlConfig.Kernel

	node.Quad = *PostProcessing.NewQuad()

	return Scene.UnmarshalChildren(value, node)
}
