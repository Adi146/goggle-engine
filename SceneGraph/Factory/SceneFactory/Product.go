package SceneFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

type Product struct {
	Scene.IScene
}

type tmpProduct struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var tmpProduct tmpProduct
	if err := value.Decode(&tmpProduct); err != nil {
		return err
	}

	sceneType, ok := typeFactory[tmpProduct.Type]
	if !ok {
		return fmt.Errorf("scene type %s is not in factory", tmpProduct.Type)
	}

	scene := reflect.New(sceneType).Interface().(Scene.IScene)

	if err := tmpProduct.Config.Decode(scene); err != nil {
		return err
	}

	product.IScene = scene
	return nil
}
