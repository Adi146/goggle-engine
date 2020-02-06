package LightNode

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const SpotLightNodeFactoryName = "Node.LightNode.SpotLightNode"

func init() {
	YamlFactory.NodeFactory[SpotLightNodeFactoryName] = reflect.TypeOf((*SpotLightNodeConfig)(nil)).Elem()
}

type SpotLightNodeConfig struct {
	Scene.ChildNodeBaseConfig
	Light.SpotLight  `yaml:"spotLight"`
	InitialDirection *Vector.Vector3 `yaml:"initialDirection"`
}

func (config SpotLightNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsChildNode()
}

func (config SpotLightNodeConfig) CreateAsChildNode() (Scene.IChildNode, error) {
	nodeBase, err := config.ChildNodeBaseConfig.CreateAsChildNode()
	if err != nil {
		return nil, err
	}

	node := &SpotLightNode{
		SpotLightNodeConfig: &config,
		IChildNode:          nodeBase,
	}

	return node, nil
}

type SpotLightNode struct {
	*SpotLightNodeConfig
	Scene.IChildNode
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.IChildNode.Tick(timeDelta)

	node.SpotLight.Position = *node.GetGlobalPosition()
	node.SpotLight.Direction = *node.GetGlobalTransformation().Inverse().Transpose().MulVector(node.InitialDirection).Normalize()

	return err
}

func (node *SpotLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(&node.SpotLight)
	}
	return nil
}
