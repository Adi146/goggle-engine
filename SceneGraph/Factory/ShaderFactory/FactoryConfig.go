package ShaderFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type FactoryConfig struct {
	ShaderConfig   map[string]FactoryEntry `yaml:"shaders"`
	DecodedShaders map[string]Shader.IShaderProgram
}

func (config *FactoryConfig) Get(name string) (Shader.IShaderProgram, error) {
	if shader, ok := config.DecodedShaders[name]; ok {
		return shader, nil
	}

	shaderConfig, ok := config.ShaderConfig[name]
	if !ok {
		return nil, fmt.Errorf("shader %s is not configured", name)
	}

	shader, err := shaderConfig.Unmarshal()
	if err == nil {
		config.DecodedShaders[name] = shader
	}

	return shader, err
}
