package CameraNode

import (
	"reflect"
	"unsafe"

	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"

	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"

	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
)

const CameraUBONodeFactoryName = "Node.CameraUBO"
const CameraUBOFactoryName = "camera"

func init() {
	YamlFactory.NodeFactory[CameraUBONodeFactoryName] = reflect.TypeOf((*CameraUBONodeConfig)(nil)).Elem()
	UniformBufferFactory.AddType(CameraUBOFactoryName, reflect.TypeOf((*Camera.UniformBuffer)(nil)).Elem())
}

type CameraUBONodeConfig struct {
	CameraNodeConfig `yaml:",inline"`
	UBO              string `yaml:"uniformBuffer"`
}

func (config *CameraUBONodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.CameraNodeConfig.CreateAsCameraNode()
	if err != nil {
		return nil, err
	}

	ubo, err := UniformBufferFactory.Get(config.UBO)
	if err != nil {
		return nil, err
	}

	ubo.UpdateData(&nodeBase.GetProjectionMatrix()[0][0], 0, int(unsafe.Sizeof(*nodeBase.GetProjectionMatrix())))

	return &CameraUBONode{
		CameraNode:     nodeBase,
		IUniformBuffer: ubo,
	}, nil
}

type CameraUBONode struct {
	*CameraNode
	UniformBuffer.IUniformBuffer
}

func (node *CameraUBONode) Tick(timeDelta float32) error {
	err := node.CameraNode.Tick(timeDelta)

	node.IUniformBuffer.UpdateData(&node.GetViewMatrix()[0][0], int(unsafe.Sizeof(*node.GetProjectionMatrix())), int(unsafe.Sizeof(*node.GetViewMatrix())))

	return err
}

func (node *CameraUBONode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.CameraPosition = node.GetGlobalPosition()
	}

	return nil
}
