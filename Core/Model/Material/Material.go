package Material

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
)

type Material struct {
	DiffuseBaseColor  GeometryMath.Vector4 `yaml:"diffuseBaseColor"`
	SpecularBaseColor GeometryMath.Vector3 `yaml:"specularBaseColor"`
	EmissiveBaseColor GeometryMath.Vector3 `yaml:"emissiveBaseColor"`

	Shininess float32 `yaml:"shininess"`

	Textures struct {
		Diffuse  *Texture.Texture2D `yaml:"diffuse"`
		Specular *Texture.Texture2D `yaml:"specular"`
		Emissive *Texture.Texture2D `yaml:"emissive"`
		Normal   *Texture.Texture2D `yaml:"normal"`
	} `yaml:"textures"`

	UvScale float32 `yaml:"uvScale"`
}

func (material *Material) Bind() error {
	if material.Textures.Diffuse != nil {
		if err := material.Textures.Diffuse.Bind(); err != nil {
			return err
		}
	}

	if material.Textures.Specular != nil {
		if err := material.Textures.Specular.Bind(); err != nil {
			return err
		}
	}

	if material.Textures.Emissive != nil {
		if err := material.Textures.Emissive.Bind(); err != nil {
			return err
		}
	}

	if material.Textures.Normal != nil {
		if err := material.Textures.Normal.Bind(); err != nil {
			return err
		}
	}

	return nil
}

func (material *Material) Unbind() {
	if material.Textures.Diffuse != nil {
		material.Textures.Diffuse.Unbind()
	}

	if material.Textures.Specular != nil {
		material.Textures.Specular.Unbind()
	}

	if material.Textures.Emissive != nil {
		material.Textures.Emissive.Unbind()
	}

	if material.Textures.Normal != nil {
		material.Textures.Normal.Unbind()
	}
}

func (material *Material) UnmarshalYAML(value *yaml.Node) error {
	type yamlConfigType Material
	yamlConfig := (yamlConfigType)(*material)
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	if yamlConfig.UvScale == 0 {
		yamlConfig.UvScale = 1
	}

	*material = (Material)(yamlConfig)
	return nil
}
