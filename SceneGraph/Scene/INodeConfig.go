package Scene

type INodeConfig interface {
	Create() (INode, error)
}
