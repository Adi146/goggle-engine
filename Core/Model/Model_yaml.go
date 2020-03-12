package Model

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	File     string                               `yaml:"file"`
	Textures map[Texture.Type][]Texture.Texture2D `yaml:"textures"`
}

func (model *Model) UnmarshalYAML(value *yaml.Node) error {
	var importErrors Error.ErrorCollection
	var importWarnings Error.ErrorCollection

	var yamlConfig yamlConfig
	if value.Kind == yaml.ScalarNode {
		value.Decode(&yamlConfig.File)
	} else {
		value.Decode(&yamlConfig)
	}

	tmpModel, result := ImportModel(yamlConfig.File)
	importErrors.Push(&result.Errors)
	importWarnings.Push(&result.Warnings)
	if result.Success() {
		for textureType, textures := range yamlConfig.Textures {
			for _, texture := range textures {
				texture.Type = textureType
				for _, mesh := range tmpModel.Meshes {
					mesh.Textures = append(mesh.Textures, &texture)
				}
			}
		}

		Log.Warn(&importWarnings, "import warnings")
		*model = *tmpModel
	}

	return importErrors.Err()
}
