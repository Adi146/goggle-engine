package YamlFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"gopkg.in/yaml.v3"
	"reflect"
)

var SceneFactory = map[string]reflect.Type{
	"sceneGraph":     reflect.TypeOf((*SceneGraphConfig)(nil)).Elem(),
	"postProcessing": reflect.TypeOf((*PostProcessingSceneConfig)(nil)).Elem(),
}

type ScenesConfig struct {
	SceneConfig   map[string]SceneConfig `yaml:"scenes"`
	DecodedScenes map[string]Scene.IScene
}

func (config *ScenesConfig) Get(name string) (Scene.IScene, error) {
	if scene, ok := config.DecodedScenes[name]; ok {
		return scene, nil
	}

	sceneConfig, ok := config.SceneConfig[name]
	if !ok {
		return nil, fmt.Errorf("scene %s is not configured", name)
	}

	scene, err := sceneConfig.Unmarshal()
	if err == nil {
		config.DecodedScenes[name] = scene
	}

	return scene, err
}

func (config *ScenesConfig) GetScenes() []Scene.IScene {
	var scenes []Scene.IScene

	for _, scene := range config.DecodedScenes {
		scenes = append(scenes, scene)
	}
	return scenes
}

type SceneConfig struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

func (config *SceneConfig) Unmarshal() (Scene.IScene, error) {
	sceneType, ok := SceneFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("scene type %s is not in factory", config.Type)
	}

	scene := reflect.New(sceneType).Interface().(Scene.IScene)

	config.Config.Decode(scene)

	if err := scene.Init(); err != nil {
		return nil, err
	}

	return scene, nil
}
