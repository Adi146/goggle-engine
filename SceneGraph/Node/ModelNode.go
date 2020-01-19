package Node

import (
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"os"
	"reflect"
)

type ModelNode struct {
	Scene.IChildNode
	*Model.Model

	File string `yaml:"file"`
}

func init() {
	Factory.NodeFactory["Node.ModelNode"] = reflect.TypeOf((*ModelNode)(nil)).Elem()
}

func (node *ModelNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	if node.Model == nil {
		file, err := os.Open(node.File)
		if err != nil {
			return err
		}
		defer file.Close()

		if node.Model, err = Model.NewModelFromFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (node *ModelNode) Tick(timeDelta float32) {
	node.ModelMatrix = node.GetGlobalTransformation()
}

func (node *ModelNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindModel(node.Model)
		for _, geometry := range node.Geometries {
			scene.GetActiveShaderProgram().BindMaterial(geometry.Material)
			geometry.Draw()
		}
	}
}
