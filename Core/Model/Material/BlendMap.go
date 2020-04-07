package Material

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
)

type BlendMaterial struct {
	BlendTexture Texture.Texture2D `yaml:"blendMap"`
	Materials    [4]Material       `yaml:"materials"`
}

func (blendMap *BlendMaterial) Unbind() {
	blendMap.BlendTexture.Unbind()
	for _, material := range blendMap.Materials {
		material.Unbind()
	}
}

func (blendMap *BlendMaterial) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		BlendTexture Texture.Texture2D `yaml:"blendMap"`
		Default      Material          `yaml:"default"`
		Red          Material          `yaml:"red"`
		Green        Material          `yaml:"green"`
		Blue         Material          `yaml:"blue"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	*blendMap = BlendMaterial{
		BlendTexture: yamlConfig.BlendTexture,
		Materials: [4]Material{
			yamlConfig.Default,
			yamlConfig.Red,
			yamlConfig.Green,
			yamlConfig.Blue,
		},
	}
	return nil
}
