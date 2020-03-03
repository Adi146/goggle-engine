package YamlFactory

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
)

type PostProcessingSceneConfig struct {
	PostProcessing.Scene `yaml:",inline"`
	ShaderConfig         ShaderFactory.Config `yaml:"shaderName"`
}

func (config *PostProcessingSceneConfig) Init() error {
	config.Scene.Shader = config.ShaderConfig

	return config.Scene.Init()
}
