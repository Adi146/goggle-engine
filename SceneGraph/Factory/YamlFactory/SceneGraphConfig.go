package YamlFactory

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

type SceneGraphConfig struct {
	Scene.Scene
	Root NodeFactory.FactoryEntry `yaml:"root"`
}

func (config *SceneGraphConfig) Init() error {
	if err := config.Scene.Init(); err != nil {
		return err
	}

	root, err := config.Root.Unmarshal("Root")
	if err != nil {
		return err
	}

	config.SetRoot(root)

	return nil
}
