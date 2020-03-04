package SceneFactory

import (
	sceneBase "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
)

type SceneGraphProduct struct {
	sceneBase.IScene
}

type tmpSceneGraphProduct struct {
	Scene.Scene
	Root NodeFactory.Product `yaml:"root"`
}

func (product *SceneGraphProduct) UnmarshalYAML(value *yaml.Node) error {
	var tmpConfig tmpSceneGraphProduct
	if err := value.Decode(&tmpConfig); err != nil {
		return err
	}

	tmpConfig.SetRoot(tmpConfig.Root)

	product.IScene = &tmpConfig.Scene
	return nil
}
