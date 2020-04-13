package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const SpotLightNodeFactoryName = "Node.LightNode.SpotLightNode"

func init() {
	Scene.NodeFactory.AddType(SpotLightNodeFactoryName, reflect.TypeOf((*SpotLightNode)(nil)).Elem())
}

type SpotLightNode struct {
	Scene.INode
	Light.UBOSpotLight
	InitDirection GeometryMath.Vector3
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.UBOSpotLight.Position.Set(*node.GetGlobalPosition())
	node.UBOSpotLight.Direction.Set(*node.GetGlobalTransformation().Inverse().Transpose().MulVector(&node.InitDirection).Normalize())
	node.UBOSpotLight.UpdateViewProjection()

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *SpotLightNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *SpotLightNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	if err := value.Decode(&node.UBOSpotLight); err != nil {
		return err
	}

	if node.InitDirection == (GeometryMath.Vector3{}) {
		node.InitDirection = node.UBOSpotLight.Direction.Get()
	}

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
