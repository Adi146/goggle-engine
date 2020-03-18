package SceneFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Scene"
	sceneGraph "github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

var (
	typeFactory = map[string]reflect.Type{
		"sceneGraph": reflect.TypeOf((*sceneGraph.Scene)(nil)).Elem(),
	}
	globalConfig FactoryConfig
)

func Get(key string) (Scene.IScene, error) {
	scene, ok := globalConfig.Scenes[key]
	if !ok {
		return nil, fmt.Errorf("scene with name %s is not configured", key)
	}

	return scene.IScene, nil
}

func SetConfig(config FactoryConfig) {
	globalConfig = config
}

func GetAll() []Scene.IScene {
	var scenes []Scene.IScene

	for _, scene := range globalConfig.Scenes {
		scenes = append(scenes, scene)
	}

	return scenes
}
