package LightNode

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"

func init() {
	YamlFactory.NodeFactory[DirectionalLightNodeFactoryName] = reflect.TypeOf((*DirectionalLightNodeConfig)(nil)).Elem()
}

type DirectionalLightNodeConfig struct {
	Scene.ChildNodeBaseConfig
	Light.DirectionalLight `yaml:"directionalLight"`
	InitialDirection       *Vector.Vector3 `yaml:"initialDirection,flow"`
}

func (config DirectionalLightNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsChildNode()
}

func (config DirectionalLightNodeConfig) CreateAsChildNode() (Scene.IChildNode, error) {
	nodeBase, err := config.ChildNodeBaseConfig.CreateAsChildNode()
	if err != nil {
		return nil, err
	}

	node := &DirectionalLightNode{
		DirectionalLightNodeConfig: &config,
		IChildNode:                 nodeBase,
	}

	if config.InitialDirection == nil {
		node.InitialDirection = &Vector.Vector3{0, 0, 1}
	}

	return node, nil
}

type DirectionalLightNode struct {
	*DirectionalLightNodeConfig
	Scene.IChildNode
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	node.DirectionalLight.Direction = *node.GetGlobalTransformation().MulVector(node.InitialDirection).Normalize()
	return nil
}

func (node *DirectionalLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(&node.DirectionalLight)
	}
	return nil
}
