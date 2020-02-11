package YamlFactory

import (
	"fmt"

	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
)

var ShaderFactory = map[string]func([]string, []string) (Shader.IShaderProgram, error){
	//"basic": BasicShader.NewBasicIShaderProgram,
	"phong":          PhongShader.NewPhongIShaderProgram,
	"postProcessing": PostProcessing.NewIShaderProgram,
}

type ShadersConfig struct {
	ShaderConfig   map[string]YamlShaderConfig `yaml:"shaders"`
	DecodedShaders map[string]Shader.IShaderProgram
	BaseConfig     *config
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
	UniformBuffers  []string `yaml:"uniformBuffers"`
}

func (config *YamlShaderConfig) Unmarshal() (Shader.IShaderProgram, error) {
	shaderConstructor, ok := ShaderFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("shader type %s is not in factory", config.Type)
	}

	shader, err := shaderConstructor(config.VertexShaders, config.FragmentShaders)
	if err != nil {
		return nil, err
	}

	for _, fboName := range config.UniformBuffers {
		fbo, err := UniformBufferFactory.Get(fboName)
		if err != nil {
			return nil, err
		}
		shader.BindObject(fbo)
	}

	return shader, nil
}
