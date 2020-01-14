package Node

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

type CameraNode struct {
	Scene.IChildNode
	*Camera.Camera

	InitialFront *Vector.Vector3 `yaml:"initialFront"`
	InitialUp    *Vector.Vector3 `yaml:"initialUp"`

	PerspectiveMatrixConfig *Factory.PerspectiveConfig `yaml:"perspective"`
	OrthogonalMatrixConfig  *Factory.OrthogonalConfig  `yaml:"orthogonal"`
}

func init() {
	Factory.NodeFactory["Node.Camera"] = reflect.TypeOf((*CameraNode)(nil)).Elem()
}

func (node *CameraNode) Init() error {
	if node.IChildNode == nil {
		node.IChildNode = Scene.NewChildNodeBase()
	}

	if node.Camera == nil {
		if node.PerspectiveMatrixConfig != nil {
			node.Camera = Camera.NewCamera(node.PerspectiveMatrixConfig.Decode())
		} else if node.OrthogonalMatrixConfig != nil {
			node.Camera = Camera.NewCamera(node.OrthogonalMatrixConfig.Decode())
		} else {
			return fmt.Errorf("no projection matrix specified")
		}
	}

	return nil
}

func (node *CameraNode) Tick(timeDelta float32) {
	node.Camera.Position = node.GetGlobalPosition()

	globalTransformation := node.GetGlobalTransformation()
	globalTransformation[0][3], globalTransformation[1][3], globalTransformation[2][3] = 0, 0, 0

	node.Camera.Front = globalTransformation.MulVector(node.InitialFront).Normalize()
	node.Camera.Up = globalTransformation.MulVector(node.InitialUp).Normalize()

	node.Camera.Tick(timeDelta)
}

func (node *CameraNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindCamera(node)
	}
}
