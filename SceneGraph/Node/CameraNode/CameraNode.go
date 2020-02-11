package CameraNode

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
	Scene.NodeConfig
	InitialFront            *Vector.Vector3                `yaml:"initialFront"`
	InitialUp               *Vector.Vector3                `yaml:"initialUp"`
	PerspectiveMatrixConfig *YamlFactory.PerspectiveConfig `yaml:"perspective"`
	OrthogonalMatrixConfig  *YamlFactory.OrthogonalConfig  `yaml:"orthogonal"`
}

func (config *CameraNodeConfig) Create() (Scene.INode, error) {
	return config.CreateAsCameraNode()
}

func (config *CameraNodeConfig) CreateAsCameraNode() (*CameraNode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &CameraNode{
		INode:  nodeBase,
		Config: config,
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
	Scene.INode
	*Camera.Camera

	Config *CameraNodeConfig
}

func (node *CameraNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.Camera.Position = node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()

	node.Camera.Front = invTransGlobalTransformation.MulVector(node.Config.InitialFront).Normalize()
	node.Camera.Up = invTransGlobalTransformation.MulVector(node.Config.InitialUp).Normalize()

	node.Camera.Tick(timeDelta)

	return err
}

func (node *CameraNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.PreRenderObjects = append(scene.PreRenderObjects, node.Camera)
		scene.CameraPosition = node.GetGlobalPosition()
	}

	return nil
}
