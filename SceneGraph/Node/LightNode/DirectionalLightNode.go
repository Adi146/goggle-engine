package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type DirectionalLightNode struct {
	Scene.IChildNode
	Light.DirectionalLight `yaml:"directionalLight"`

	InitialDirection *Vector.Vector3 `yaml:"initialDirection,flow"`
}

func init() {
	YamlFactory.NodeFactory["Node.LightNode.DirectionalLightNode"] = reflect.TypeOf((*DirectionalLightNode)(nil)).Elem()
}

func (node *DirectionalLightNode) Init(nodeID string) error {
	if node.IChildNode == nil {
		node.IChildNode = &Scene.ChildNodeBase{}
		if err := node.IChildNode.Init(nodeID); err != nil {
			return err
		}
	}

	if node.InitialDirection == nil {
		node.InitialDirection = &Vector.Vector3{0, 0, 1}
	}

	return nil
}

func (node *DirectionalLightNode) Tick(timeDelta float32) error {
	node.DirectionalLight.Direction = *node.GetGlobalTransformation().MulVector(node.InitialDirection).Normalize()
	return nil
}

func (node *DirectionalLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(&node.DirectionalLight)
	}
	return nil
}
