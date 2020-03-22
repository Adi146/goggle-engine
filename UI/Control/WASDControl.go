package Control

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const WASDControlFactoryName = "UI.Control.WASDControl"

func init() {
	Scene.NodeFactory.AddType(WASDControlFactoryName, reflect.TypeOf((*WASDControl)(nil)).Elem())
}

type WASDControl struct {
	Scene.INode

	KeyboardSensitivity float32
	MouseSensitivity    float32

	yaw   float32
	pitch float32
}

func (node *WASDControl) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	scene := node.GetScene()
	if scene != nil {
		xRel, yRel := scene.Window.GetMouseInput().GetRelativeMovement()
		node.Rotate(GeometryMath.Radians(xRel*node.MouseSensitivity), GeometryMath.Radians(yRel*node.MouseSensitivity))

		if scene.Window.GetKeyboardInput().IsKeyPressed("W") {
			node.MoveForwards(node.KeyboardSensitivity * timeDelta)
		}

		if scene.Window.GetKeyboardInput().IsKeyPressed("S") {
			node.MoveForwards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.Window.GetKeyboardInput().IsKeyPressed("A") {
			node.MoveSidewards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.Window.GetKeyboardInput().IsKeyPressed("D") {
			node.MoveSidewards(node.KeyboardSensitivity * timeDelta)
		}
	}

	return err
}

func (node *WASDControl) MoveForwards(amount float32) {
	vec := GeometryMath.Vector3{0, 0, 1}
	node.Translate(vec.MulScalar(-amount))
}

func (node *WASDControl) MoveSidewards(amount float32) {
	vec := GeometryMath.Vector3{1, 0, 0}
	node.Translate(vec.MulScalar(amount))
}

func (node *WASDControl) Translate(vec *GeometryMath.Vector3) {
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(GeometryMath.Translate(vec)))
}

func (node *WASDControl) Rotate(x float32, y float32) {
	currentPosition := node.GetLocalPosition()

	node.yaw -= x
	node.pitch -= y

	node.SetLocalTransformation(GeometryMath.Translate(currentPosition))
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(GeometryMath.RotateY(node.yaw)))
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(GeometryMath.RotateX(node.pitch)))
}

func (node *WASDControl) UnmarshalYAML(value *yaml.Node) error {
	if node.INode == nil {
		node.INode = &Scene.Node{}
	}
	if err := value.Decode(node.INode); err != nil {
		return err
	}

	yamlConfig := struct {
		KeyboardSensitivity float32 `yaml:"keyboardSensitivity"`
		MouseSensitivity    float32 `yaml:"mouseSensitivity"`
	}{
		KeyboardSensitivity: node.KeyboardSensitivity,
		MouseSensitivity:    node.MouseSensitivity,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.KeyboardSensitivity = yamlConfig.KeyboardSensitivity
	node.MouseSensitivity = yamlConfig.MouseSensitivity

	return nil
}
