package Scene

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"

type NodeConfig struct {
}

func (config *NodeConfig) Create() (INode, error) {
	return &Node{
		transformation: Matrix.Identity(),
		Config:         config,
	}, nil
}
