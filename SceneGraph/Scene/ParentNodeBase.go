package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type ParentNodeBase struct {
	*NodeBase
	children []IChildNode
}

func (node *ParentNodeBase) Init(nodeID string) error {
	if node.NodeBase == nil {
		node.NodeBase = &NodeBase{
			scene:          nil,
			transformation: Matrix.Identity(),
		}
		if err := node.NodeBase.Init(nodeID); err != nil {
			return err
		}
	}

	return nil
}

func (node *ParentNodeBase) setScene(scene *Scene) {
	node.NodeBase.setScene(scene)

	for _, child := range node.children {
		child.setScene(scene)
	}
}

func (node *ParentNodeBase) AddChild(child IChildNode) {
	node.children = append(node.children, child)
	child.setParent(node)
}

func (node *ParentNodeBase) GetChildren() []IChildNode {
	return node.children
}

func (node *ParentNodeBase) TickChildren(timeDelta float32) error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(Error.NewErrorWithFields(child.Tick(timeDelta), child.GetLogFields()))
		if childAsParent, isParent := child.(IParentNode); isParent {
			err.Push(Error.NewErrorWithFields(childAsParent.TickChildren(timeDelta), childAsParent.GetLogFields()))
		}
	}

	return err.Err()
}

func (node *ParentNodeBase) DrawChildren() error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Draw())
		if childAsParent, isParent := child.(IParentNode); isParent {
			err.Push(childAsParent.DrawChildren())
		}
	}

	return err.Err()
}
