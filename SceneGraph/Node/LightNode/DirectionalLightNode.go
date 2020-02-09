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
	Scene.NodeConfig
	Light.DirectionalLight `yaml:"directionalLight"`
	InitialDirection       *Vector.Vector3 `yaml:"initialDirection,flow"`
}

func (config *DirectionalLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &DirectionalLightNode{
		INode:  nodeBase,
		Config: config,
	}

	if config.InitialDirection == nil {
		config.InitialDirection = &Vector.Vector3{0, 0, 1}
	}

	return node, nil
}

type DirectionalLightNode struct {
	Scene.INode
	Config *DirectionalLightNodeConfig
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.Config.DirectionalLight.Direction = *node.GetGlobalTransformation().MulVector(node.Config.InitialDirection).Normalize()

	return err
}

func (node *DirectionalLightNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.PreRenderObjects = append(scene.PreRenderObjects, &node.Config.DirectionalLight)
	}

	return nil
}
