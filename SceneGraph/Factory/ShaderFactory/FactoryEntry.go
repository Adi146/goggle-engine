package ShaderFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
)

type FactoryEntry struct {
	Type            string   `yaml:"type"`
	VertexShaders   []string `yaml:"vertexShaders"`
	FragmentShaders []string `yaml:"fragmentShaders"`
	UniformBuffers  []string `yaml:"uniformBuffers"`
}

func (config *FactoryEntry) Unmarshal() (Shader.IShaderProgram, error) {
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
