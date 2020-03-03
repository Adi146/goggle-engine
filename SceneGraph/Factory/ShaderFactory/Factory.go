package ShaderFactory

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
)

var (
	typeFactory = map[string]func([]string, []string) (Shader.IShaderProgram, error){
		"phong":          PhongShader.NewPhongIShaderProgram,
		"postProcessing": PostProcessing.NewIShaderProgram,
	}
	globalConfig FactoryConfig
)

func AddType(key string, constructor func([]string, []string) (Shader.IShaderProgram, error)) {
	typeFactory[key] = constructor
}

func Get(key string) (Shader.IShaderProgram, error) {
	return globalConfig.Get(key)
}

func SetConfig(config FactoryConfig) {
	globalConfig = config
}
