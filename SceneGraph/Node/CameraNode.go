package Node

import (
	"fmt"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const CameraNodeFactoryName = "Node.Camera"

func init() {
	YamlFactory.NodeFactory[CameraNodeFactoryName] = reflect.TypeOf((*CameraNodeConfig)(nil)).Elem()
}

type CameraNodeConfig struct {
	Scene.ChildNodeBaseConfig
	InitialFront            *Vector.Vector3                `yaml:"initialFront"`
	InitialUp               *Vector.Vector3                `yaml:"initialUp"`
	PerspectiveMatrixConfig *YamlFactory.PerspectiveConfig `yaml:"perspective"`
	OrthogonalMatrixConfig  *YamlFactory.OrthogonalConfig  `yaml:"orthogonal"`
}

func (config CameraNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsChildNode()
}

func (config CameraNodeConfig) CreateAsChildNode() (Scene.IChildNode, error) {
	nodeBase, err := config.ChildNodeBaseConfig.CreateAsChildNode()
	if err != nil {
		return nil, err
	}

	node := &CameraNode{
		CameraNodeConfig: &config,
		IChildNode:       nodeBase,
	}

	if config.PerspectiveMatrixConfig != nil {
		node.Camera = Camera.NewCamera(config.PerspectiveMatrixConfig.Decode())
	} else if config.OrthogonalMatrixConfig != nil {
		node.Camera = Camera.NewCamera(config.OrthogonalMatrixConfig.Decode())
	} else {
		return nil, fmt.Errorf("no projection matrix specified")
	}

	if config.InitialUp == nil {
		config.InitialUp = &Vector.Vector3{0, 1, 0}
	}

	if config.InitialFront == nil {
		config.InitialFront = &Vector.Vector3{0, 0, 1}
	}

	return node, nil
}

type CameraNode struct {
	*CameraNodeConfig
	Scene.IChildNode
	*Camera.Camera
}

func (node *CameraNode) Tick(timeDelta float32) error {
	err := node.IChildNode.Tick(timeDelta)

	node.Camera.Position = node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()

	node.Camera.Front = invTransGlobalTransformation.MulVector(node.InitialFront).Normalize()
	node.Camera.Up = invTransGlobalTransformation.MulVector(node.InitialUp).Normalize()

	node.Camera.Tick(timeDelta)

	return err
}

func (node *CameraNode) Draw() error {
	scene := node.GetScene()
	if scene != nil && scene.GetActiveShaderProgram() != nil {
		return scene.GetActiveShaderProgram().BindObject(node.Camera)
	}
	return nil
}
