package CameraNode

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/MatrixFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const CameraNodeFactoryName = "Node.Camera"
const CameraUBOFactoryName = "camera"

func init() {
	NodeFactory.AddType(CameraNodeFactoryName, reflect.TypeOf((*CameraNodeConfig)(nil)).Elem())
	UniformBufferFactory.AddType(CameraUBOFactoryName, reflect.TypeOf((*Camera.UniformBuffer)(nil)).Elem())
}

type CameraNodeConfig struct {
	Scene.NodeConfig
	FrontVector      Vector.Vector3             `yaml:"front"`
	UpVector         Vector.Vector3             `yaml:"up"`
	ProjectionMatrix MatrixFactory.MatrixConfig `yaml:"projection"`
	UBO              string                     `yaml:"uniformBuffer"`
}

func (config *CameraNodeConfig) Create() (Scene.INode, error) {
	config.SetDefaults()

	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	ubo, err := UniformBufferFactory.Get(config.UBO)
	if err != nil {
		return nil, err
	}

	projectionMatrix := config.ProjectionMatrix

	cameraBase := Camera.NewCamera(projectionMatrix.Matrix4x4)
	camera := ubo.(Camera.ICamera)
	camera.Set(*cameraBase)

	node := &CameraNode{
		INode:   nodeBase,
		ICamera: camera,

		Config: config,
	}

	return node, nil
}

func (config *CameraNodeConfig) SetDefaults() {
	if config.UpVector.Length() == 0 {
		config.UpVector = Vector.Vector3{0, 1, 0}
	}

	if config.FrontVector.Length() == 0 {
		config.FrontVector = Vector.Vector3{0, 0, 1}
	}
}

type CameraNode struct {
	Scene.INode
	Camera.ICamera

	Config *CameraNodeConfig
}

func (node *CameraNode) Tick(timeDelta float32) error {
	position := node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()
	front := invTransGlobalTransformation.MulVector(&node.Config.FrontVector).Normalize()
	up := invTransGlobalTransformation.MulVector(&node.Config.UpVector).Normalize()

	node.ICamera.SetViewMatrix(*Matrix.LookAt(position, position.Add(front), up))

	return nil
}

func (node *CameraNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.CameraPosition = node.GetGlobalPosition()
	}

	return nil
}
