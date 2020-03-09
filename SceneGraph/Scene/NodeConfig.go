package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type NodeConfig struct {
}

func (config *NodeConfig) Create() (INode, error) {
	return &Node{
		transformation: GeometryMath.Identity(),
		Config:         config,
	}, nil
}
