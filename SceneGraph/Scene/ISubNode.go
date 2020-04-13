package Scene

type ISubNode interface {
	INode
	SetBase(node INode)
}
