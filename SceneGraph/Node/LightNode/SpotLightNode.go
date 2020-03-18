package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const SpotLightNodeFactoryName = "Node.LightNode.SpotLightNode"

func init() {
	NodeFactory.AddType(SpotLightNodeFactoryName, reflect.TypeOf((*SpotLightNodeConfig)(nil)).Elem())
}

type SpotLightNodeConfig struct {
	Scene.NodeConfig
	Light.SpotLight `yaml:"spotLight"`
	UBOSpotLight    Light.UBOSpotLight `yaml:",inline"`
}

func (config *SpotLightNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &SpotLightNode{
		INode:      nodeBase,
		ISpotLight: &config.UBOSpotLight,
		Config:     config,
	}

	return node, nil
}

type SpotLightNode struct {
	Scene.INode
	Light.ISpotLight

	Config *SpotLightNodeConfig
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetPosition(*node.GetGlobalPosition())
	node.SetDirection(*node.GetGlobalTransformation().Inverse().Transpose().MulVector(&node.Config.SpotLight.Direction).Normalize())

	return err
}
