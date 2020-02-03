package YamlFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

type SceneGraphConfig struct {
	Scene.Scene
	Root NodeConfig `yaml:"root"`
}

func (config *SceneGraphConfig) Init() error {
	if err := config.Scene.Init(); err != nil {
		return err
	}

	root, err := config.Root.Unmarshal("Root")
	if err != nil {
		return err
	}

	if rootAsParent, isParent := root.(Scene.IParentNode); isParent {
		config.SetRoot(rootAsParent)
	} else {
		return fmt.Errorf("root is not a parent node")
	}

	return nil
}
