package Material

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
)

const (
	BlendMap Texture.Type = "blendMap"
)

type BlendMaterial struct {
	BlendMap  *Texture.Texture2D `yaml:"blendMap"`
	Materials [4]Material        `yaml:"materials"`
}

func (blendMap *BlendMaterial) Unbind() {
	blendMap.BlendMap.Unbind()
	for _, material := range blendMap.Materials {
		material.Unbind()
	}
}

func (blendMap *BlendMaterial) SetWrapMode(mode Texture.WrapMode) {
	for _, material := range blendMap.Materials {
		material.SetWrapMode(mode)
	}
}

func (blendMap *BlendMaterial) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		BlendMap *Texture.Texture2D `yaml:"blendMap"`
		Default  Material           `yaml:"default"`
		Red      Material           `yaml:"r"`
		Green    Material           `yaml:"g"`
		Blue     Material           `yaml:"b"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	if yamlConfig.BlendMap != nil {
		yamlConfig.BlendMap.Type = BlendMap
	}

	*blendMap = BlendMaterial{
		BlendMap: yamlConfig.BlendMap,
		Materials: [4]Material{
			yamlConfig.Default,
			yamlConfig.Red,
			yamlConfig.Green,
			yamlConfig.Blue,
		},
	}
	return nil
}
