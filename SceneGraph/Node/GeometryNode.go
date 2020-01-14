package Node

import (
	"github.com/Adi146/goggle-engine/Core/Geometry"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"os"
	"reflect"
)

type GeometryNode struct {
	Scene.IChildNode
	*Geometry.Geometry

	File string `yaml:"file"`
}

func init() {
	Factory.NodeFactory["Node.GeometryNode"] = reflect.TypeOf((*GeometryNode)(nil)).Elem()
}

func (node *GeometryNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	if node.Geometry == nil {
		file, err := os.Open(node.File)
		if err != nil {
			return err
		}
		defer file.Close()

		if node.Geometry, err = Geometry.NewGeometryFromFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (node *GeometryNode) Tick(timeDelta float32) {
	node.Geometry.ModelMatrix = node.GetGlobalTransformation()
}

func (node *GeometryNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindGeometry(node.Geometry)
		node.Geometry.Draw()
	}
}
