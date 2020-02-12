package ShaderFactory

import (
	"fmt"

	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
)

var (
	typeFactory = map[string]func([]string, []string) (Shader.IShaderProgram, error){
		"phong":          PhongShader.NewPhongIShaderProgram,
		"postProcessing": PostProcessing.NewIShaderProgram,
	}
	globalConfig ShadersConfig
)

func AddType(key string, constructor func([]string, []string) (Shader.IShaderProgram, error)) {
	typeFactory[key] = constructor
}

func Get(key string) (Shader.IShaderProgram, error) {
	return globalConfig.Get(key)
}

func SetConfig(config ShadersConfig) {
	globalConfig = config
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
	UniformBuffers  []string `yaml:"uniformBuffers"`
}

func (config *YamlShaderConfig) Unmarshal() (Shader.IShaderProgram, error) {
	shaderConstructor, ok := typeFactory[config.Type]
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
		if err := shader.BindObject(fbo); err != nil {
			return nil, err
		}
	}

	return shader, nil
}
