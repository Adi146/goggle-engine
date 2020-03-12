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
	"reflect"
)

const (
	PostProcessingNodeFactoryName   = "Node.PostProcessing"
	PostProcessingShaderFactoryName = "postProcessing"
	OffScreenFramebufferName        = "offScreen"
)

func init() {
	NodeFactory.AddType(PostProcessingNodeFactoryName, reflect.TypeOf((*PostProcessingNodeConfig)(nil)).Elem())
	ShaderFactory.AddType(PostProcessingShaderFactoryName, PostProcessing.NewIShaderProgram)
	FrameBufferFactory.AddType(OffScreenFramebufferName, reflect.TypeOf((*FrameBuffer.OffScreenBuffer)(nil)).Elem())
}

type PostProcessingNodeConfig struct {
	Scene.NodeConfig
	Shader      ShaderFactory.Config      `yaml:"shader"`
	FrameBuffer FrameBufferFactory.Config `yaml:"frameBuffer"`
	Kernel      *PostProcessing.Kernel    `yaml:",inline"`
}

func (config *PostProcessingNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	quad, err := PostProcessing.NewQuad()
	if err != nil {
		return nil, err
	}

	node := &PostProcessingNode{
		INode:  nodeBase,
		Config: config,
		quad:   quad,
	}

	return node, nil
}

type PostProcessingNode struct {
	Scene.INode
	Config *PostProcessingNodeConfig
	quad   *Model.Mesh
}

func (node *PostProcessingNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	if node.Config.Kernel != nil {
		if err := node.Config.Shader.BindObject(node.Config.Kernel); err != nil {
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

	node.Config.FrameBuffer.Bind()
	node.Config.FrameBuffer.Clear()
	err := scene.Draw(shader, node, scene)
	node.Config.FrameBuffer.Unbind()
	if err != nil {
		return err
	}

	scene.Clear()

	node.Config.Shader.Bind()
	if shader != nil {
		defer shader.Bind()
	}
	return node.quad.Draw(node.Config.Shader, node, scene)
}
