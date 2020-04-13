package CameraNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
	"reflect"

	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const CameraNodeFactoryName = "Node.Camera"

func init() {
	Scene.NodeFactory.AddType(CameraNodeFactoryName, reflect.TypeOf((*CameraNode)(nil)).Elem())
}

type CameraNode struct {
	Scene.INode

	FrontVector GeometryMath.Vector3
	UpVector    GeometryMath.Vector3
}

func (node *CameraNode) Tick(timeDelta float32) error {
	position := node.GetGlobalPosition()

	invTransGlobalTransformation := node.GetGlobalTransformation().Inverse().Transpose()
	front := invTransGlobalTransformation.MulVector(&node.FrontVector).Normalize()
	up := invTransGlobalTransformation.MulVector(&node.UpVector).Normalize()

	if scene := node.GetScene(); scene != nil {
		scene.Camera.Update(*position, *front, *up)
	}

	return nil
}

func (node *CameraNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *CameraNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	yamlConfig := struct {
		FrontVector GeometryMath.Vector3 `yaml:"front"`
		UpVector    GeometryMath.Vector3 `yaml:"up"`
	}{
		FrontVector: node.FrontVector,
		UpVector:    node.UpVector,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.FrontVector = yamlConfig.FrontVector
	node.UpVector = yamlConfig.UpVector

	if node.FrontVector == (GeometryMath.Vector3{}) {
		node.FrontVector = GeometryMath.Vector3{0, 0, 1}
	}

	if node.UpVector == (GeometryMath.Vector3{}) {
		node.UpVector = GeometryMath.Vector3{0, 1, 0}
	}

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
