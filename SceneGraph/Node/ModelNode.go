package Node

import (
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
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
		model, err := Model.ImportModel(node.File)
		if err != nil {
			return err
		}
		node.Model = model
	}

	return nil
}

func (node *ModelNode) Tick(timeDelta float32) {
	node.ModelMatrix = node.GetGlobalTransformation()
}

func (node *ModelNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindObject(node.Model)
		for _, mesh := range node.Meshes {
			scene.GetActiveShaderProgram().BindObject(mesh.Material)
			mesh.Draw()
		}
	}
}
