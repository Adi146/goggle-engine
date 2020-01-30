package YamlFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
)

var ShaderFactory = map[string]func([]string, []string) (Shader.IShaderProgram, error){
	//"basic": BasicShader.NewBasicIShaderProgram,
	"phong": PhongShader.NewPhongIShaderProgram,
}

type ShadersConfig struct {
	ShaderConfig   map[string]YamlShaderConfig `yaml:"shaders"`
	DecodedShaders map[string]Shader.IShaderProgram
}

func (config *ShadersConfig) Get(name string) (Shader.IShaderProgram, error) {
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
