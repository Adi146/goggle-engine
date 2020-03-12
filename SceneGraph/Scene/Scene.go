package Scene

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Utils/Log"
)

type Scene struct {
	coreScene.SceneBase
	Root INode
}

func (scene *Scene) Init() error {
	return nil
}

func (scene *Scene) Tick(timeDelta float32) {
	scene.SceneBase.Tick(timeDelta)

	if scene.Root != nil {
		Log.Error(scene.Root.Tick(timeDelta), "tick error")
	}
}

func (scene *Scene) SetRoot(node INode) {
	node.setScene(scene)
	scene.Root = node
}
