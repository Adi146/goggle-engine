package Scene

import "github.com/Adi146/goggle-engine/Utils"

type ParentNodeBase struct {
	*NodeBase
	children []IChildNode
}

func NewParentNodeBase() *ParentNodeBase {
	return &ParentNodeBase{
		NodeBase: NewNodeBase(),
		children: []IChildNode{},
	}
}

func (node *ParentNodeBase) Init() error {
	if node.NodeBase == nil {
		node.NodeBase = NewNodeBase()
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
	var err Utils.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Tick(timeDelta))
		if childAsParent, isParent := child.(IParentNode); isParent {
			err.Push(childAsParent.TickChildren(timeDelta))
		}
	}

	return err.Err()
}

func (node *ParentNodeBase) DrawChildren() error {
	var err Utils.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Draw())
		if childAsParent, isParent := child.(IParentNode); isParent {
			err.Push(childAsParent.DrawChildren())
		}
	}

	return err.Err()
}
