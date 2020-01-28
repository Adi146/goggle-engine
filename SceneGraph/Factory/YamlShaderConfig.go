package Factory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type YamlShaderConfig struct {
	Type            string   `yaml:"type"`
	VertexShaders   []string `yaml:"vertexShaders"`
	FragmentShaders []string `yaml:"fragmentShaders"`
}

func (config *YamlShaderConfig) Unmarshal() (Shader.IShaderProgram, error) {
	shaderConstructor, ok := ShaderFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("shader type %s is not in factory", config.Type)
	}

	return shaderConstructor(config.VertexShaders, config.FragmentShaders)
}
