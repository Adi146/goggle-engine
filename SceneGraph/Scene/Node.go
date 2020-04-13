package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"gopkg.in/yaml.v3"
	"reflect"
)

const (
	NodeFactoryName = "Scene.Node"
)

func init() {
	NodeFactory.AddType(NodeFactoryName, reflect.TypeOf((*Node)(nil)).Elem())
}

type Node struct {
	scene          *Scene
	Transformation *GeometryMath.Matrix4x4

	children []INode
	parent   INode
}

func (node *Node) AddChild(child INode) {
	node.children = append(node.children, child)
}

func (node *Node) GetChildren() []INode {
	return node.children
}

func (node *Node) GetParent() INode {
	return node.parent
}

func (node *Node) SetParent(parent INode) {
	node.parent = parent
	if parent != nil {
		node.SetScene(parent.GetScene())
	} else {
		node.SetScene(nil)
	}
}

func (node *Node) GetScene() *Scene {
	return node.scene
}

func (node *Node) SetScene(scene *Scene) {
	node.scene = scene

	for _, child := range node.children {
		child.SetScene(scene)
	}
}

func (node *Node) GetLocalTransformation() *GeometryMath.Matrix4x4 {
	return node.Transformation
}

func (node *Node) SetLocalTransformation(matrix *GeometryMath.Matrix4x4) {
	node.Transformation = matrix
}

func (node *Node) GetLocalRotation() []GeometryMath.EulerAngles {
	return GeometryMath.ExtractFromMatrix(node.GetLocalTransformation())
}

func (node *Node) GetLocalPosition() *GeometryMath.Vector3 {
	return node.GetLocalTransformation().MulVector(&GeometryMath.Vector3{0, 0, 0})
}

func (node *Node) GetGlobalTransformation() *GeometryMath.Matrix4x4 {
	if parent := node.GetParent(); parent == nil {
		return node.GetLocalTransformation()
	} else {
		return node.parent.GetGlobalTransformation().Mul(node.GetLocalTransformation())
	}
}

func (node *Node) GetGlobalRotation() []GeometryMath.EulerAngles {
	return GeometryMath.ExtractFromMatrix(node.GetGlobalTransformation())
}

func (node *Node) GetGlobalPosition() *GeometryMath.Vector3 {
	return node.GetGlobalTransformation().MulVector(&GeometryMath.Vector3{0, 0, 0})
}

func (node *Node) Tick(timeDelta float32) error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Tick(timeDelta))
	}

	return err.Err()
}

func (node *Node) GetBase() INode {
	return node
}

func (node *Node) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Transformation []GeometryMath.Matrix4x4 `yaml:"transformation"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	if len(yamlConfig.Transformation) >= 1 || node.Transformation == nil || *node.Transformation == (GeometryMath.Matrix4x4{}) {
		node.SetLocalTransformation(GeometryMath.Identity())
		for _, transformation := range yamlConfig.Transformation {
			node.SetLocalTransformation(node.GetLocalTransformation().Mul(&transformation))
		}
	}

	return nil
}
