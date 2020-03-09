package Control

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const WASDControlFactoryName = "UI.Control.WASDControl"

func init() {
	NodeFactory.AddType(WASDControlFactoryName, reflect.TypeOf((*WASDControlConfig)(nil)).Elem())
}

type WASDControlConfig struct {
	Scene.NodeConfig
	KeyboardSensitivity float32 `yaml:"keyboardSensitivity"`
	MouseSensitivity    float32 `yaml:"mouseSensitivity"`
}

func (config *WASDControlConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()

	return &WASDControl{
		INode:  nodeBase,
		Config: config,
	}, err
}

type WASDControl struct {
	Scene.INode

	Config *WASDControlConfig

	yaw   float32
	pitch float32
}

func (node *WASDControl) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	scene := node.GetScene()
	if scene != nil {
		xRel, yRel := scene.GetMouseInput().GetRelativeMovement()
		node.Rotate(GeometryMath.Radians(xRel*node.Config.MouseSensitivity), GeometryMath.Radians(yRel*node.Config.MouseSensitivity))

		if scene.GetKeyboardInput().IsKeyPressed("W") {
			node.MoveForwards(node.Config.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("S") {
			node.MoveForwards(-node.Config.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("A") {
			node.MoveSidewards(-node.Config.KeyboardSensitivity * timeDelta)
		}

		if scene.GetKeyboardInput().IsKeyPressed("D") {
			node.MoveSidewards(node.Config.KeyboardSensitivity * timeDelta)
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
