package Material

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
)

type Material struct {
	DiffuseBaseColor  GeometryMath.Vector4
	SpecularBaseColor GeometryMath.Vector3
	EmissiveBaseColor GeometryMath.Vector3

	Shininess float32

	Textures []Texture.ITexture
}

func (material *Material) Bind() {
	for _, texture := range material.Textures {
		texture.Bind()
	}
}

func (material *Material) Unbind() {
	for _, texture := range material.Textures {
		texture.Unbind()
	}
}

func (material *Material) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		DiffuseBaseColor  GeometryMath.Vector4 `yaml:"diffuseBaseColor"`
		SpecularBaseColor GeometryMath.Vector3 `yaml:"specularBaseColor"`
		EmissiveBaseColor GeometryMath.Vector3 `yaml:"emissiveBaseColor"`

		Shininess float32 `yaml:"shininess"`

		Textures map[Texture.Type][]Texture.Texture2D `yaml:"textures"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	material.DiffuseBaseColor = yamlConfig.DiffuseBaseColor
	material.SpecularBaseColor = yamlConfig.SpecularBaseColor
	material.EmissiveBaseColor = yamlConfig.EmissiveBaseColor

	material.Shininess = yamlConfig.Shininess

	for textureType, textures := range yamlConfig.Textures {
		for _, texture := range textures {
			texture.Type = textureType
			material.Textures = append(material.Textures, &texture)
		}
	}

	return nil
}
