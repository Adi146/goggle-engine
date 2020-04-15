package ModelNode

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const ModelSlaveNodeFactoryName = "Node.ModelNode.ModelSlaveNode"

func init() {
	Scene.NodeFactory.AddType(ModelSlaveNodeFactoryName, reflect.TypeOf((*ModelSlaveNode)(nil)).Elem())
}

type ModelSlaveNode struct {
	Scene.INode
	Mesh.IMesh
	Master    *ModelNode
	MasterID  string
}

func (node *ModelSlaveNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	if scene := node.GetScene(); scene != nil {
		nodeByID := scene.Root.GetGrandChildById(node.MasterID)
		newMaster, isMaster := nodeByID.(*ModelNode)
		if !isMaster {
			return fmt.Errorf("node with id %s is no ModelNode (actual type %T)", node.MasterID, nodeByID)
		}

		if node.Master != newMaster {
			if err := newMaster.AddSlave(node); err != nil {
				return err
			}
		}
	}

	if node.IMesh != nil {
		node.IMesh.SetModelMatrix(*node.GetGlobalTransformation())
	}

	return err
}

func (node *ModelSlaveNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *ModelSlaveNode) UnmarshalYAML(value *yaml.Node) error {
	if err := Scene.UnmarshalBase(value, node); err != nil {
		return err
	}

	yamlConfig := struct {
		MasterID string `yaml:"masterID"`
	}{
		MasterID: node.MasterID,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.MasterID = yamlConfig.MasterID

	return Scene.UnmarshalChildren(value, node, Scene.NodeFactoryName)
}
