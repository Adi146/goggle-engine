package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
)

const (
	NodeFactoryName = "Scene.Node"
)

func init() {
	NodeFactory.AddType(NodeFactoryName, reflect.TypeOf((*Node)(nil)).Elem())
}

type Node struct {
	scene          *Scene
	Transformation GeometryMath.Matrix4x4

	children map[string]INode
	parent   INode

	id string

	events map[string]IEvent
}

func (node *Node) AddChild(child INode, id string) {
	if node.children == nil {
		node.children = map[string]INode{}
	}

	node.children[id] = child
}

func (node *Node) GetChildren() map[string]INode {
	return node.children
}

func (node *Node) GetGrandChildById(id string) INode {
	if id == "" {
		return node
	}

	split := strings.SplitN(id, ".", 2)
	child := node.GetChildren()[split[0]]
	if child != nil {
		if len(split) == 1 {
			return child
		} else {
			return child.GetGrandChildById(split[1])
		}
	}

	return nil
}

func (node *Node) GetID() string {
	return node.id
}

func (node *Node) SetID(id string) {
	node.id = id
}

func (node *Node) GetParent() INode {
	return node.parent
}

func (node *Node) SetParent(parent INode, childID string) {
	node.parent = parent
	if parent != nil {
		node.SetScene(parent.GetScene())
		parentID := parent.GetID()
		if parentID == "" {
			node.SetID(childID)
		} else {
			node.SetID(parentID + "." + childID)
		}
	} else {
		node.SetScene(nil)
		node.SetID("")
	}
}

func (node *Node) GetScene() *Scene {
	return node.scene
}

func (node *Node) SetScene(scene *Scene) {
	node.scene = scene

	for _, child := range node.GetChildren() {
		child.SetScene(scene)
	}
}

func (node *Node) GetLocalTransformation() GeometryMath.Matrix4x4 {
	return node.Transformation
}

func (node *Node) SetLocalTransformation(matrix GeometryMath.Matrix4x4) {
	node.Transformation = matrix
}

func (node *Node) GetLocalRotation() []GeometryMath.EulerAngles {
	return GeometryMath.ExtractFromMatrix(node.GetLocalTransformation())
}

func (node *Node) GetLocalPosition() GeometryMath.Vector3 {
	return node.GetLocalTransformation().MulVector(GeometryMath.Vector3{0, 0, 0})
}

func (node *Node) GetGlobalTransformation() GeometryMath.Matrix4x4 {
	if parent := node.GetParent(); parent == nil {
		return node.GetLocalTransformation()
	} else {
		return node.parent.GetGlobalTransformation().Mul(node.GetLocalTransformation())
	}
}

func (node *Node) GetGlobalRotation() []GeometryMath.EulerAngles {
	return GeometryMath.ExtractFromMatrix(node.GetGlobalTransformation())
}

func (node *Node) GetGlobalPosition() GeometryMath.Vector3 {
	return node.GetGlobalTransformation().MulVector(GeometryMath.Vector3{0, 0, 0})
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

func (node *Node) Start() error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Start())
	}

	return err.Err()
}

func (node *Node) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Transformation GeometryMath.Matrix4x4 `yaml:"transformation"`
	}{
		Transformation: node.Transformation,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}
	node.SetLocalTransformation(yamlConfig.Transformation)
	if node.GetLocalTransformation() == (GeometryMath.Matrix4x4{}) {
		node.SetLocalTransformation(GeometryMath.Identity())
	}

	return nil
}

func (node *Node) AddEvent(event IEvent, id string) {
	if node.events == nil {
		node.events = map[string]IEvent{}
	}

	node.events[id] = event
}

func (node *Node) GetEventByID(id string) IEvent {
	return node.events[id]
}
