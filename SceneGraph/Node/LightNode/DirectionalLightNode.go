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

	return nil
}

func (node *DirectionalLightNode) Tick(timeDelta float32) {
	node.DirectionalLight.Direction = *node.GetGlobalTransformation().MulVector(node.InitialDirection).Normalize()
}

func (node *DirectionalLightNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindDirectionalLight(&node.DirectionalLight)
	}
}
