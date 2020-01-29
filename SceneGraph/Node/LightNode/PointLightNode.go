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

func (node *PointLightNode) Init(nodeID string) error {
	if node.IChildNode == nil {
		node.IChildNode = &Scene.ChildNodeBase{}
		if err := node.IChildNode.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *PointLightNode) Tick(timeDelta float32) error {
	node.PointLight.Position = *node.GetGlobalPosition()
	return nil
}

func (node *PointLightNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetFrameBuffer() != nil {
		return scene.GetFrameBuffer().GetShaderProgram().BindObject(&node.PointLight)
	}
	return nil
}
