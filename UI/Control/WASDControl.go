package Control

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const WASDControlFactoryName = "UI.Control.WASDControl"

func init() {
	YamlFactory.NodeFactory[WASDControlFactoryName] = reflect.TypeOf((*WASDControlConfig)(nil)).Elem()
}

type WASDControlConfig struct {
	Scene.IntermediateNodeBaseConfig
	KeyboardSensitivity float32 `yaml:"keyboardSensitivity"`
	MouseSensitivity    float32 `yaml:"mouseSensitivity"`
}

func (config WASDControlConfig) Create() (Scene.INode, error) {
	return config.CreateAsIntermediateNode()
}

func (config WASDControlConfig) CreateAsIntermediateNode() (Scene.IIntermediateNode, error) {
	nodeBase, err := config.IntermediateNodeBaseConfig.CreateAsIntermediateNode()
	if err != nil {
		return nil, err
	}

	node := &WASDControl{
		WASDControlConfig: &config,
		IIntermediateNode: nodeBase,
	}

	return node, nil
}

type WASDControl struct {
	*WASDControlConfig
	Scene.IIntermediateNode

	yaw   float32
	pitch float32
}

func (node *WASDControl) Tick(timeDelta float32) error {
	err := node.IIntermediateNode.Tick(timeDelta)

	scene := node.GetScene()
	if scene != nil {
		xRel, yRel := scene.GetMouseInput().GetRelativeMovement()
		node.Rotate(Angle.Radians(xRel*node.MouseSensitivity), Angle.Radians(yRel*node.MouseSensitivity))

		if scene.GetKeyboardInput().IsKeyPressed("W") {
			node.MoveForwards(node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("S") {
			node.MoveForwards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("A") {
			node.MoveSidewards(-node.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("D") {
			node.MoveSidewards(node.KeyboardSensitivity * timeDelta)
		}
	}

	return err
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
