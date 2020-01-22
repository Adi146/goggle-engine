package Control

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type WASDControl struct {
	Scene.IIntermediateNode

	KeyboardSensitivity float32 `yaml:"keyboardSensitivity"`
	MouseSensitivity    float32 `yaml:"mouseSensitivity"`

	yaw   float32
	pitch float32
}

func init() {
	Factory.NodeFactory["UI.Control.WASDControl"] = reflect.TypeOf((*WASDControl)(nil)).Elem()
}

func (node *WASDControl) Init(nodeID string) error {
	if node.IIntermediateNode == nil {
		node.IIntermediateNode = &Scene.IntermediateNodeBase{}
		if err := node.IIntermediateNode.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *WASDControl) Tick(timeDelta float32) error {
	scene := node.GetScene()
	if scene != nil {
		xRel, yRel := scene.GetWindow().GetMouseInput().GetRelativeMovement()
		node.Rotate(Angle.Radians(xRel*node.MouseSensitivity), Angle.Radians(yRel*node.MouseSensitivity))

		if scene.GetWindow().GetKeyboardInput().IsKeyPressed("W") {
			node.MoveForwards(node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetWindow().GetKeyboardInput().IsKeyPressed("S") {
			node.MoveForwards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetWindow().GetKeyboardInput().IsKeyPressed("A") {
			node.MoveSidewards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetWindow().GetKeyboardInput().IsKeyPressed("D") {
			node.MoveSidewards(node.KeyboardSensitivity * timeDelta)
		}
	}

	return nil
}

func (node *WASDControl) MoveForwards(amount float32) {
	vec := Vector.Vector3{0, 0, 1}
	node.Translate(vec.MulScalar(-amount))
}

func (node *WASDControl) MoveSidewards(amount float32) {
	vec := Vector.Vector3{1, 0, 0}
	node.Translate(vec.MulScalar(amount))
}

func (node *WASDControl) Translate(vec *Vector.Vector3) {
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.Translate(vec)))
}

func (node *WASDControl) Rotate(x float32, y float32) {
	currentPosition := node.GetLocalPosition()

	node.yaw -= x
	node.pitch -= y

	node.SetLocalTransformation(Matrix.Translate(currentPosition))
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.RotateY(node.yaw)))
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.RotateX(node.pitch)))
}
