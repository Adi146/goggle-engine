package Node

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

const CameraNodeFactoryName = "Node.Camera"

type CameraNode struct {
	Scene.IChildNode
	*Camera.Camera

	InitialFront *Vector.Vector3 `yaml:"initialFront"`
	InitialUp    *Vector.Vector3 `yaml:"initialUp"`

	PerspectiveMatrixConfig *Factory.PerspectiveConfig `yaml:"perspective"`
	OrthogonalMatrixConfig  *Factory.OrthogonalConfig  `yaml:"orthogonal"`
}

func init() {
	Factory.NodeFactory[CameraNodeFactoryName] = reflect.TypeOf((*CameraNode)(nil)).Elem()
}

func (node *CameraNode) Init(nodeID string) error {
	if node.IChildNode == nil {
		node.IChildNode = &Scene.ChildNodeBase{}
		if err := node.IChildNode.Init(nodeID); err != nil {
			return err
		}
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

func (node *CameraNode) Tick(timeDelta float32) error {
	node.Camera.Position = node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()

	node.Camera.Front = invTransGlobalTransformation.MulVector(node.InitialFront).Normalize()
	node.Camera.Up = invTransGlobalTransformation.MulVector(node.InitialUp).Normalize()

	node.Camera.Tick(timeDelta)

	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(node.Camera)
	}

	return nil
}
