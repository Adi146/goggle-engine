package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const (
	DirectionalLightNodeFactoryName = "Node.LightNode.DirectionalLightNode"
)

func init() {
	Scene.NodeFactory.AddType(DirectionalLightNodeFactoryName, reflect.TypeOf((*DirectionalLightNode)(nil)).Elem())
}

type DirectionalLightNode struct {
	Scene.INode
	Light.UBODirectionalLight
	InitDirection GeometryMath.Vector3
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.UBODirectionalLight.SetDirection(*node.GetGlobalTransformation().MulVector(&node.InitDirection).Normalize())

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *DirectionalLightNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *DirectionalLightNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	if err := value.Decode(&node.UBODirectionalLight); err != nil {
		return err
	}

	if node.InitDirection == (GeometryMath.Vector3{}) {
		node.InitDirection = node.UBODirectionalLight.Direction.Get()
	}

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
