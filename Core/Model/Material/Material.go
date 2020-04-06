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

	DiffuseTexture  *Texture.Texture2D
	SpecularTexture *Texture.Texture2D
	EmissiveTexture *Texture.Texture2D
	NormalTexture   *Texture.Texture2D
}

func (material *Material) Bind() error {
	if material.DiffuseTexture != nil {
		if err := material.DiffuseTexture.Bind(); err != nil {
			return err
		}
	}

	if material.SpecularTexture != nil {
		if err := material.SpecularTexture.Bind(); err != nil {
			return err
		}
	}

	if material.EmissiveTexture != nil {
		if err := material.EmissiveTexture.Bind(); err != nil {
			return err
		}
	}

	if material.NormalTexture != nil {
		if err := material.NormalTexture.Bind(); err != nil {
			return err
		}
	}

	return nil
}

func (material *Material) Unbind() {
	if material.DiffuseTexture != nil {
		material.DiffuseTexture.Unbind()
	}

	if material.SpecularTexture != nil {
		material.SpecularTexture.Unbind()
	}

	if material.EmissiveTexture != nil {
		material.EmissiveTexture.Unbind()
	}

	if material.NormalTexture != nil {
		material.NormalTexture.Unbind()
	}
}

func (material *Material) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		DiffuseBaseColor  GeometryMath.Vector4 `yaml:"diffuseBaseColor"`
		SpecularBaseColor GeometryMath.Vector3 `yaml:"specularBaseColor"`
		EmissiveBaseColor GeometryMath.Vector3 `yaml:"emissiveBaseColor"`

		Shininess float32 `yaml:"shininess"`

		Textures struct {
			DiffuseTexture  *Texture.Texture2D `yaml:"diffuse"`
			SpecularTexture *Texture.Texture2D `yaml:"specular"`
			EmissiveTexture *Texture.Texture2D `yaml:"emissive"`
			NormalTexture   *Texture.Texture2D `yaml:"normal"`
		} `yaml:"textures"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	material.DiffuseBaseColor = yamlConfig.DiffuseBaseColor
	material.SpecularBaseColor = yamlConfig.SpecularBaseColor
	material.EmissiveBaseColor = yamlConfig.EmissiveBaseColor

	material.Shininess = yamlConfig.Shininess

	material.DiffuseTexture = yamlConfig.Textures.DiffuseTexture
	material.SpecularTexture = yamlConfig.Textures.SpecularTexture
	material.EmissiveTexture = yamlConfig.Textures.EmissiveTexture
	material.NormalTexture = yamlConfig.Textures.NormalTexture

	return nil
}

func (material *Material) Merge(m2 *Material) {
	if m2.DiffuseBaseColor != (GeometryMath.Vector4{}) {
		material.DiffuseBaseColor = m2.DiffuseBaseColor
	}

	if m2.SpecularBaseColor != (GeometryMath.Vector3{}) {
		material.SpecularBaseColor = m2.SpecularBaseColor
	}

	if m2.EmissiveBaseColor != (GeometryMath.Vector3{}) {
		material.EmissiveBaseColor = m2.EmissiveBaseColor
	}

	if m2.Shininess != 0 {
		material.Shininess = m2.Shininess
	}

	if m2.DiffuseTexture != nil {
		material.DiffuseTexture = m2.DiffuseTexture
	}

	if m2.SpecularTexture != nil {
		material.SpecularTexture = m2.SpecularTexture
	}

	if m2.EmissiveTexture != nil {
		material.EmissiveTexture = m2.EmissiveTexture
	}

	if m2.NormalTexture != nil {
		material.NormalTexture = m2.NormalTexture
	}
}
