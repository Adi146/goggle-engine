package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type DirectionalLightNode struct {
	Scene.IChildNode
	Light.DirectionalLight `yaml:"directionalLight"`

	InitialDirection *Vector.Vector3 `yaml:"initialDirection,flow"`
}

func init() {
	Factory.NodeFactory["Node.LightNode.DirectionalLightNode"] = reflect.TypeOf((*DirectionalLightNode)(nil)).Elem()
}

func (node *DirectionalLightNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	if node.InitialDirection == nil {
		node.InitialDirection = &Vector.Vector3{0, 0, 1}
	}

	return nil
}

func (node *DirectionalLightNode) Tick(timeDelta float32) {
	node.DirectionalLight.Direction = *node.GetGlobalTransformation().MulVector(node.InitialDirection).Normalize()

	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindObject(&node.DirectionalLight)
	}
}
