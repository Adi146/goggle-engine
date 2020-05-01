package UtilNode

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

const (
	TriggerFactoryName = "Node.UtilNode.TriggerNode"
	TriggerEventName   = "trigger"
)

func init() {
	Scene.NodeFactory.AddType(TriggerFactoryName, reflect.TypeOf((*TriggerNode)(nil)).Elem())
}

type TriggerNode struct {
	Scene.INode
	TriggerVolume             BoundingVolume.IBoundingVolume
	TransformedBoundingVolume BoundingVolume.IBoundingVolume
	TriggerOn                 []BoundingVolume.ICollisionObject
	TriggerOnIDs              []string
	TriggerEvent              Scene.Event
}

func (node *TriggerNode) Start() error {
	if scene := node.GetScene(); scene != nil {
		for _, triggerID := range node.TriggerOnIDs {
			nodeByID := scene.Root.GetGrandChildById(triggerID)
			collisionObject, isCollisionObject := nodeByID.(BoundingVolume.ICollisionObject)
			if !isCollisionObject {
				return fmt.Errorf("trigger object %s is not a collision object", triggerID)
			}

			node.TriggerOn = append(node.TriggerOn, collisionObject)
		}
	}

	return nil
}

func (node *TriggerNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.TransformedBoundingVolume = node.TriggerVolume.Transform(node.GetGlobalTransformation())

	var triggerNodes []BoundingVolume.ICollisionObject
	for _, collisionObjection := range node.TriggerOn {
		if node.TransformedBoundingVolume.IntersectsWith(collisionObjection.GetBoundingVolume()) {
			triggerNodes = append(triggerNodes, collisionObjection)
		}
	}
	node.TriggerEvent.NotifyListeners(triggerNodes)

	return err
}

func (node *TriggerNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *TriggerNode) GetBoundingVolume() BoundingVolume.IBoundingVolume {
	return node.TransformedBoundingVolume
}

func (node *TriggerNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		TriggerVolume BoundingVolume.Ptr `yaml:"triggerVolume"`
		TriggerOn     []string           `yaml:"triggerOn"`
	}{
		TriggerVolume: BoundingVolume.Ptr{
			IBoundingVolume: node.TriggerVolume,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.TriggerVolume = yamlConfig.TriggerVolume.IBoundingVolume
	node.TriggerOnIDs = yamlConfig.TriggerOn
	node.AddEvent(&node.TriggerEvent, TriggerEventName)

	return Scene.UnmarshalChildren(value, node)
}
