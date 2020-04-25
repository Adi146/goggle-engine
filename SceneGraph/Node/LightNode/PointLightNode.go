package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const PointLightNodeFactoryName = "Node.LightNode.PointLightNode"

func init() {
	Scene.NodeFactory.AddType(PointLightNodeFactoryName, reflect.TypeOf((*PointLightNode)(nil)).Elem())
}

type PointLightNode struct {
	Scene.INode
	Light.UBOPointLight
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.UBOPointLight.PointLight.Position.Set(node.GetGlobalPosition())

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *PointLightNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *PointLightNode) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&node.UBOPointLight); err != nil {
		return err
	}

	return Scene.UnmarshalChildren(value, node)
}
