package Scene

import (
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type ParentNodeBaseConfig struct {
	NodeBaseConfig
}

func (config ParentNodeBaseConfig) Create() (INode, error) {
	return config.CreateAsParentNode()
}

func (config ParentNodeBaseConfig) CreateAsParentNode() (IParentNode, error) {
	nodeBase, err := config.NodeBaseConfig.Create()

	node := &ParentNodeBase{
		ParentNodeBaseConfig: config,
		INode:                nodeBase,
	}

	return node, err
}

type ParentNodeBase struct {
	ParentNodeBaseConfig
	INode
	children []IChildNode
}

func (node *ParentNodeBase) setScene(scene *Scene) {
	node.INode.setScene(scene)

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
	}

	return err.Err()
}

func (node *ParentNodeBase) DrawChildren() error {
	var err Error.ErrorCollection

	for _, child := range node.GetChildren() {
		err.Push(child.Draw())
	}

	return err.Err()
}
