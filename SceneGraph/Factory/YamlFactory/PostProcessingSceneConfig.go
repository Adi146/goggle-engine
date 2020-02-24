package YamlFactory

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
)

type PostProcessingSceneConfig struct {
	PostProcessing.Scene `yaml:",inline"`
	ShaderName string `yaml:"shaderName"`
}

func (config *PostProcessingSceneConfig) Init() error {
	shader, err := ShaderFactory.Get(config.ShaderName)
	if err != nil {
		return err
	}

	config.Scene.Shader = shader

	return config.Scene.Init()
}
