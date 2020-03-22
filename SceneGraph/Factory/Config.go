package Factory

import (
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Utils/Log"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"

	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type config struct {
	OpenGlLogging bool
	SceneGraph    *Scene.Scene
}

func (config *config) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		OpenGlLogging bool        `yaml:"openGlLogging"`
		SceneGraph    Scene.Scene `yaml:"scene"`
	}

	if err := value.Decode(&Shader.Factory); err != nil {
		return err
	}

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	config.OpenGlLogging = yamlConfig.OpenGlLogging
	config.SceneGraph = &yamlConfig.SceneGraph

	return nil
}

func ReadConfig(file *os.File) (coreScene.IScene, error) {
	var config config

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	if config.OpenGlLogging {
		Log.EnableDebugLogging()
	}

	return config.SceneGraph, nil
}
