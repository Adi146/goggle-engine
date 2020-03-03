package ShaderFactory

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Shader.IShaderProgram
}

func (config *Config) UnmarshalYAML(value *yaml.Node) error {
	var shaderName string
	value.Decode(&shaderName)

	shader, err := Get(shaderName)
	if err != nil {
		return err
	}

	config.IShaderProgram = shader
	return nil
}
