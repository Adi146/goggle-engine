package Scene

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

func (node *ParentNodeBase) TickChildren(timeDelta float32) {
	for _, child := range node.GetChildren() {
		child.Tick(timeDelta)
		if childAsParent, isParent := child.(IParentNode); isParent {
			childAsParent.TickChildren(timeDelta)
		}
	}
}

func (node *ParentNodeBase) DrawChildren() {
	for _, child := range node.GetChildren() {
		child.Draw()
		if childAsParent, isParent := child.(IParentNode); isParent {
			childAsParent.DrawChildren()
		}
	}
}
