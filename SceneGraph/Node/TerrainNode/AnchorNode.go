package TerrainNode

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Terrain"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

const AnchorNodeFactoryName = "Node.TerrainNode.AnchorNode"

func init() {
	Scene.NodeFactory.AddType(AnchorNodeFactoryName, reflect.TypeOf((*AnchorNode)(nil)).Elem())
}

type AnchorNode struct {
	Scene.Node
	Terrain Terrain.ITerrain
}

func (node *AnchorNode) GetGlobalTransformation() *GeometryMath.Matrix4x4 {
	if node.Terrain != nil {
		terrainHeight := node.Terrain.GetHeightAt(*node.GetLocalPosition())
		return GeometryMath.Translate(&GeometryMath.Vector3{0, terrainHeight, 0}).Mul(node.Node.GetGlobalTransformation())
	} else {
		return node.Node.GetGlobalTransformation()
	}
}

func (node *AnchorNode) GetGlobalRotation() []GeometryMath.EulerAngles {
	return GeometryMath.ExtractFromMatrix(node.GetGlobalTransformation())
}

func (node *AnchorNode) GetGlobalPosition() *GeometryMath.Vector3 {
	return node.GetGlobalTransformation().MulVector(&GeometryMath.Vector3{0, 0, 0})
}

func (node *AnchorNode) GetBase() Scene.INode {
	return node
}

func (node *AnchorNode) SetParent(parent Scene.INode) {
	node.Node.SetParent(parent)

	if asTerrain, isTerrain := parent.(Terrain.ITerrain); isTerrain {
		node.Terrain = asTerrain
	}
}
