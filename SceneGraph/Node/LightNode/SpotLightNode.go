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
	Light.SpotLight

	FrontVector GeometryMath.Vector3
	UpVector    GeometryMath.Vector3
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	position := node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()
	front := invTransGlobalTransformation.MulVector(node.FrontVector).Normalize()
	up := invTransGlobalTransformation.MulVector(node.UpVector).Normalize()

	node.SpotLight.Update(position, front, up)

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *SpotLightNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *SpotLightNode) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&node.SpotLight); err != nil {
		return err
	}

	if node.FrontVector == (GeometryMath.Vector3{}) {
		node.FrontVector = GeometryMath.Vector3{0, 0, 1}
	}

	if node.UpVector == (GeometryMath.Vector3{}) {
		node.UpVector = GeometryMath.Vector3{0, 1, 0}
	}

	return Scene.UnmarshalChildren(value, node)
}
