package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type SpotLightNode struct {
	Scene.IChildNode
	Light.SpotLight `yaml:"spotLight"`

	InitialDirection *Vector.Vector3 `yaml:"initialDirection"`
}

func init() {
	Factory.NodeFactory["Node.LightNode.SpotLightNode"] = reflect.TypeOf((*SpotLightNode)(nil)).Elem()
}

func (node *SpotLightNode) Init(nodeID string) error {
	if node.IChildNode == nil {
		node.IChildNode = &Scene.ChildNodeBase{}
		if err := node.IChildNode.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *SpotLightNode) Tick(timeDelta float32) error {
	node.SpotLight.Position = *node.GetGlobalPosition()
	node.SpotLight.Direction = *node.GetGlobalTransformation().Inverse().Transpose().MulVector(node.InitialDirection).Normalize()
	return nil
}

func (node *SpotLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetFrameBuffer().GetShaderProgram() != nil {
		return scene.GetFrameBuffer().GetShaderProgram().BindObject(&node.SpotLight)
	}
	return nil
}
