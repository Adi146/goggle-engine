package ShaderFactory

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Shader.IShaderProgram
}

func (config *Config) UnmarshalYAML(value *yaml.Node) error {
	var name string
	value.Decode(&name)

	ubo, err := Get(name)
	if err != nil {
		return err
	}

	config.IShaderProgram = ubo
	return nil
}