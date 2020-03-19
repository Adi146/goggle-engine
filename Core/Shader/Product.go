package Shader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"gopkg.in/yaml.v3"
)

type Product struct {
	shaderProgram   IShaderProgram
	VertexShaders   []string
	FragmentShaders []string
	Constructor     func([]string, []string) (IShaderProgram, error)
	FrameBuffers    []FrameBufferFactory.Config
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Type            string                      `yaml:"type"`
		VertexShaders   []string                    `yaml:"vertexShaders"`
		FragmentShaders []string                    `yaml:"fragmentShaders"`
		FrameBuffers    []FrameBufferFactory.Config `yaml:"frameBuffers"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	shaderConstructor, ok := Factory.Constructors[yamlConfig.Type]
	if !ok {
		return fmt.Errorf("shader type %s is not in factory", yamlConfig.Type)
	}

	product.VertexShaders = yamlConfig.VertexShaders
	product.FragmentShaders = yamlConfig.FragmentShaders
	product.Constructor = shaderConstructor
	product.FrameBuffers = yamlConfig.FrameBuffers

	return nil
}

func (product *Product) Get() (IShaderProgram, error) {
	var errors Error.ErrorCollection

	if product.shaderProgram == nil {
		tmpProgram, err := product.Constructor(product.VertexShaders, product.FragmentShaders)
		if err != nil {
			return nil, err
		}

		for _, fboConfig := range product.FrameBuffers {
			if err := tmpProgram.BindObject(fboConfig.IFrameBuffer); err != nil {
				errors.Push(err)
			}
		}

		product.shaderProgram = tmpProgram
	}

	return product.shaderProgram, errors.Err()
}
