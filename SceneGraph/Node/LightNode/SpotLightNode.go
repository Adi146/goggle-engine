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
	Scene.NodeConfig
	Light.SpotLight  `yaml:"spotLight"`
	InitialDirection *Vector.Vector3 `yaml:"initialDirection"`
}

func (config *SpotLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &SpotLightNode{
		INode:  nodeBase,
		Config: config,
	}

	return node, nil
}

type SpotLightNode struct {
	Scene.INode

	Config *SpotLightNodeConfig
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.Config.SpotLight.Position = *node.GetGlobalPosition()
	node.Config.SpotLight.Direction = *node.GetGlobalTransformation().Inverse().Transpose().MulVector(node.Config.InitialDirection).Normalize()

	return err
}

func (node *SpotLightNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.PreRenderObjects = append(scene.PreRenderObjects, &node.Config.SpotLight)
	}

	return nil
}
