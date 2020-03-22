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
	Light.IDirectionalLight
	InitDirection GeometryMath.Vector3
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.IDirectionalLight.SetDirection(*node.GetGlobalTransformation().MulVector(&node.InitDirection).Normalize())

	if scene := node.GetScene(); scene != nil {
		scene.AddPreRenderObject(node)
	}

	return err
}

func (node *DirectionalLightNode) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	if node.IDirectionalLight == nil {
		node.IDirectionalLight = &Light.UBODirectionalLight{}
	}
	if err := value.Decode(node.IDirectionalLight); err != nil {
		return err
	}

	if node.InitDirection == (GeometryMath.Vector3{}) {
		node.InitDirection = node.IDirectionalLight.GetDirection()
	}

	return nil
}
