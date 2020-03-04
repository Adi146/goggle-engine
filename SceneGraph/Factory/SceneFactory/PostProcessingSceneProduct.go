package SceneFactory

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"gopkg.in/yaml.v3"
)

type PostProcessingProduct struct {
	Scene.IScene
}

type tmpPostProcessingProduct struct {
	PostProcessing.Scene `yaml:",inline"`
	ShaderConfig         ShaderFactory.Config `yaml:"shaderName"`
}

func (product *PostProcessingProduct) UnmarshalYAML(value *yaml.Node) error {
	var tmpConfig tmpPostProcessingProduct
	if err := value.Decode(&tmpConfig); err != nil {
		return err
	}

	tmpConfig.Scene.Shader = tmpConfig.ShaderConfig

	product.IScene = &tmpConfig.Scene
	return nil
}
