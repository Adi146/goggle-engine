package LightNode

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type PointLightNode struct {
	Scene.IChildNode
	Light.PointLight `yaml:"pointLight"`
}

func init() {
	Factory.NodeFactory["Node.LightNode.PointLightNode"] = reflect.TypeOf((*PointLightNode)(nil)).Elem()
}

func (node *PointLightNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	return nil
}

func (node *PointLightNode) Tick(timeDelta float32) {
	node.PointLight.Position = *node.GetGlobalPosition()
}

func (node *PointLightNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindPointLight(&node.PointLight)
	}
}
