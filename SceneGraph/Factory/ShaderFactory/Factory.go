package ShaderFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

var (
	typeFactory  = map[string]func([]string, []string) (Shader.IShaderProgram, error){}
	globalConfig FactoryConfig
)

func AddType(key string, constructor func([]string, []string) (Shader.IShaderProgram, error)) {
	typeFactory[key] = constructor
}

func Get(key string) (Shader.IShaderProgram, error) {
	shader, ok := globalConfig.Shaders[key]
	if !ok {
		return nil, fmt.Errorf("shader with name %s is not configured", key)
	}

	return shader.IShaderProgram, nil
}

func SetConfig(config FactoryConfig) {
	globalConfig = config
}
