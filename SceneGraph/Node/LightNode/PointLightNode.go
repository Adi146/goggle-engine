package LightNode

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const PointLightNodeFactoryName = "Node.LightNode.PointLightNode"

func init() {
	YamlFactory.NodeFactory[PointLightNodeFactoryName] = reflect.TypeOf((*PointLightNodeConfig)(nil)).Elem()
}

type PointLightNodeConfig struct {
	Scene.ChildNodeBaseConfig
	Light.PointLight `yaml:"pointLight"`
}

func (config PointLightNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsChildNode()
}

func (config PointLightNodeConfig) CreateAsChildNode() (Scene.IChildNode, error) {
	nodeBase, err := config.ChildNodeBaseConfig.CreateAsChildNode()
	if err != nil {
		return nil, err
	}

	node := &PointLightNode{
		PointLightNodeConfig: &config,
		IChildNode:           nodeBase,
	}

	return node, nil
}

type PointLightNode struct {
	*PointLightNodeConfig
	Scene.IChildNode
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	node.PointLight.Position = *node.GetGlobalPosition()
	return nil
}

func (node *PointLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(&node.PointLight)
	}
	return nil
}
