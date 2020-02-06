package SceneGraph

import (
	"reflect"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

const RotorFactoryName = "SceneGraph.Rotor"

func init() {
	YamlFactory.NodeFactory[RotorFactoryName] = reflect.TypeOf((*RotorConfig)(nil)).Elem()
}

type RotorConfig struct {
	Scene.IntermediateNodeBaseConfig
	Speed float32 `yaml:"speed"`
}

func (config RotorConfig) Create() (Scene.INode, error) {
	return config.CreateAsIntermediateNode()
}

func (config RotorConfig) CreateAsIntermediateNode() (Scene.IIntermediateNode, error) {
	nodeBase, err := config.IntermediateNodeBaseConfig.CreateAsIntermediateNode()
	if err != nil {
		return nil, err
	}

	node := &Rotor{
		RotorConfig:       &config,
		IIntermediateNode: nodeBase,
	}

	return node, err
}

type Rotor struct {
	*RotorConfig
	Scene.IIntermediateNode
}

func (node *Rotor) Tick(timeDelta float32) error {
	node.SetLocalTransformation(node.GetLocalTransformation().Mul(Matrix.RotateY(node.Speed * timeDelta)))

	return nil
}
