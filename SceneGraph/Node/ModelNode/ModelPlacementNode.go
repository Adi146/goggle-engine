package ModelNode

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Terrain"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
	"math/rand"
	"reflect"
	"time"
)

const ModelPlacementNodeFactoryName = "Node.ModelNode.ModelPlacementNode"

func init() {
	Scene.NodeFactory.AddType(ModelPlacementNodeFactoryName, reflect.TypeOf((*ModelPlacementNode)(nil)).Elem())
}

type ModelPlacementNode struct {
	ModelNode
	Terrain.PlacementMap
	RandomGenerator *rand.Rand
	AddToParent     bool
}

func (node *ModelPlacementNode) PlaceModels() error {
	nodeID := node.GetID()
	boundingVolume := node.GetBoundingVolume()

	offsetX := float32(node.NumColumns)/2 - 0.5
	offsetZ := float32(node.NumRows)/2 - 0.5

	type newSlaveConfig struct {
		Transformation GeometryMath.Matrix4x4
		BoundingVolume BoundingVolume.IBoundingVolume
	}
	var newSlaveConfigs []newSlaveConfig

	for z := 0; z < node.NumRows; z += 1 {
	NextPosition:
		for x := 0; x < node.NumColumns; x += 1 {
			var probability float32
			switch v := boundingVolume.(type) {
			case BoundingVolume.AABB:
				probability = node.GetProbabilityAtArea(
					x+int(GeometryMath.Floor(v.Min.X())),
					z+int(GeometryMath.Floor(v.Min.Z())),
					x+int(GeometryMath.Ceil(v.Max.X())),
					z+int(GeometryMath.Ceil(v.Max.Z())),
				)
			case BoundingVolume.Sphere:
				probability = node.GetProbabilityAtArea(
					x+int(GeometryMath.Floor(v.Center.X()-v.Radius)),
					z+int(GeometryMath.Floor(v.Center.Z()-v.Radius)),
					x+int(GeometryMath.Ceil(v.Center.X()+v.Radius)),
					z+int(GeometryMath.Ceil(v.Center.Z()+v.Radius)),
				)
			}

			if probability >= float32(node.RandomGenerator.NormFloat64()*0.125+0.5) {
				slaveLocalMatrix := GeometryMath.Translate(GeometryMath.Vector3{float32(x) - offsetX, 0, float32(z) - offsetZ})
				slaveBoundingBox := boundingVolume.Transform(slaveLocalMatrix)

				if slaveBoundingBox.IntersectsWith(boundingVolume) {
					continue NextPosition
				}

				for _, child := range node.GetChildren() {
					if collisionObject, isCollisionObject := child.(BoundingVolume.ICollisionObject); isCollisionObject {
						if slaveBoundingBox.IntersectsWith(collisionObject.GetBoundingVolume()) {
							continue NextPosition
						}
					}
				}

				for _, newSlave := range newSlaveConfigs {
					if slaveBoundingBox.IntersectsWith(newSlave.BoundingVolume) {
						continue NextPosition
					}
				}

				newSlaveConfigs = append(newSlaveConfigs, newSlaveConfig{
					Transformation: slaveLocalMatrix,
					BoundingVolume: slaveBoundingBox,
				})
			}
		}
	}

	Log.Info(fmt.Sprintf("%d new slaves found", len(newSlaveConfigs)))

	newSlaves := make([]*ModelSlaveNode, len(newSlaveConfigs))
	for i, newSlaveConfig := range newSlaveConfigs {
		newSlaves[i] = &ModelSlaveNode{
			INode: &Scene.Node{
				Transformation: newSlaveConfig.Transformation,
			},
			IMesh:    nil,
			Master:   nil,
			MasterID: nodeID,
		}
	}

	if err := node.AddSlave(newSlaves...); err != nil {
		return err
	}
	node.AddSlavesToSceneGraph()

	return nil
}

func (node *ModelPlacementNode) SetParent(parent Scene.INode, childID string) {
	node.INode.SetParent(parent, childID)
	if node.AddToParent {
		node.AddSlavesToSceneGraph()
	}
}

func (node *ModelPlacementNode) AddSlavesToSceneGraph() {
	var slaveParent Scene.INode
	if node.AddToParent {
		if node.GetParent() == nil {
			return
		}
		slaveParent = node.GetParent()
	} else {
		slaveParent = node
	}

	for i, slave := range node.Slaves {
		Scene.AddChild(slaveParent, slave, fmt.Sprintf("slave_%d", i))
	}
}

func (node *ModelPlacementNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		PlacementMap Terrain.PlacementMap `yaml:"placementMap"`
		ModelNode    ModelNode            `yaml:"model"`
		Seed         int64                `yaml:"seed"`
		AddToParent  bool                 `yaml:"addToParent"`
	}{
		Seed:      time.Now().UnixNano(),
		ModelNode: node.ModelNode,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.PlacementMap = yamlConfig.PlacementMap
	node.ModelNode = yamlConfig.ModelNode
	node.RandomGenerator = rand.New(rand.NewSource(yamlConfig.Seed))
	node.AddToParent = yamlConfig.AddToParent

	if err := node.PlaceModels(); err != nil {
		return err
	}

	return Scene.UnmarshalChildren(value, node)
}
