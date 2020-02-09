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
	Scene.NodeConfig
	Light.PointLight `yaml:"pointLight"`
}

func (config *PointLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &PointLightNode{
		INode:  nodeBase,
		Config: config,
	}

	return node, nil
}

type PointLightNode struct {
	Scene.INode

	Config *PointLightNodeConfig
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.Config.PointLight.Position = *node.GetGlobalPosition()

	return err
}

func (node *PointLightNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.PreRenderObjects = append(scene.PreRenderObjects, &node.Config.PointLight)
	}

	return nil
}
