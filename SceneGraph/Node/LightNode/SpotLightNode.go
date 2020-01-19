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

func (node *SpotLightNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	return nil
}

func (node *SpotLightNode) Tick(timeDelta float32) {
	node.SpotLight.Position = *node.GetGlobalPosition()
	node.SpotLight.Direction = *node.GetGlobalTransformation().Inverse().Transpose().MulVector(node.InitialDirection).Normalize()

	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindObject(&node.SpotLight)
	}
}
