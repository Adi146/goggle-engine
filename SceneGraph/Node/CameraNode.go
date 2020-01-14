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

	if node.InitialUp == nil {
		node.InitialUp = &Vector.Vector3{0, 1, 0}
	}

	if node.InitialFront == nil {
		node.InitialFront = &Vector.Vector3{0, 0, 1}
	}

	return nil
}

func (node *CameraNode) Tick(timeDelta float32) {
	node.Camera.Position = node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()

	node.Camera.Front = invTransGlobalTransformation.MulVector(node.InitialFront).Normalize()
	node.Camera.Up = invTransGlobalTransformation.MulVector(node.InitialUp).Normalize()

	node.Camera.Tick(timeDelta)
}

func (node *CameraNode) Draw() {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		scene.GetActiveShaderProgram().BindCamera(node)
	}
}
