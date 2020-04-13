package Scene

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
)

type Scene struct {
	coreScene.SceneBase
	Root INode
}

func (scene *Scene) Tick(timeDelta float32) {
	scene.SceneBase.Tick(timeDelta)
	if scene.Root != nil {
		Log.Error(scene.Root.Tick(timeDelta), "tick error")
	}
}

func (scene *Scene) SetRoot(node INode) {
	node.SetScene(scene)
	scene.Root = node
}

func (scene *Scene) UnmarshalYAML(value *yaml.Node) error {
	if scene.Root == nil {
		scene.SetRoot(&Node{})
	}

	if err := value.Decode(&scene.SceneBase); err != nil {
		return err
	}

	if err := value.Decode(scene.Root); err != nil {
		return err
	}

	return UnmarshalChildren(value, scene.Root, NodeFactoryName)
}
