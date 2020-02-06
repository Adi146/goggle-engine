package Scene

type IParentNodeConfig interface {
	INodeConfig

	CreateAsParentNode() (IParentNode, error)
}

type IParentNode interface {
	INode

	AddChild(child IChildNode)
	GetChildren() []IChildNode

	TickChildren(timeDelta float32) error
	DrawChildren() error
}
