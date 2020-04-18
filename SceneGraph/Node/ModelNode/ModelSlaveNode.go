package ModelNode

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const ModelSlaveNodeFactoryName = "Node.ModelNode.ModelSlaveNode"

func init() {
	Scene.NodeFactory.AddType(ModelSlaveNodeFactoryName, reflect.TypeOf((*ModelSlaveNode)(nil)).Elem())
}

type IMaster interface {
	Scene.INode
	AddSlave(node ...*ModelSlaveNode) error
}

type slaves []*ModelSlaveNode

type ModelSlaveNode struct {
	Scene.INode
	Mesh.IMesh
	Master   IMaster
	MasterID string
}

func (node *ModelSlaveNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	if scene := node.GetScene(); scene != nil {
		if node.Master == nil || node.Master.GetID() != node.MasterID {
			nodeByID := scene.Root.GetGrandChildById(node.MasterID)
			newMaster, isMaster := nodeByID.(IMaster)
			if !isMaster {
				return fmt.Errorf("node with id %s is no IMaster (actual type %T)", node.MasterID, nodeByID)
			}

			if err := newMaster.AddSlave(node); err != nil {
				return err
			}
		}
	}

	if node.IMesh != nil {
		node.IMesh.SetModelMatrix(node.GetGlobalTransformation())
	}

	return err
}

func (node *ModelSlaveNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *ModelSlaveNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		MasterID string `yaml:"masterID"`
	}{
		MasterID: node.MasterID,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.MasterID = yamlConfig.MasterID

	return Scene.UnmarshalChildren(value, node)
}

func (slaves slaves) GetGlobalTransformations() []GeometryMath.Matrix4x4 {
	transformations := make([]GeometryMath.Matrix4x4, len(slaves))
	for i := range slaves {
		transformations[i] = slaves[i].GetGlobalTransformation()
	}
	return transformations
}

func (slaves slaves) SetInstancedMeshes(master IMaster, instancedMeshes ...*Mesh.InstancedMesh) {
	for i := range slaves {
		slaves[i].IMesh = instancedMeshes[i]
		slaves[i].Master = master
	}
}
