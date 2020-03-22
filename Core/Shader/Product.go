package Shader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"gopkg.in/yaml.v3"
)

type Product struct {
	shaderProgram   IShaderProgram
	VertexShaders   []string
	FragmentShaders []string
	Constructor     func([]string, []string) (IShaderProgram, error)
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Type            string   `yaml:"type"`
		VertexShaders   []string `yaml:"vertexShaders"`
		FragmentShaders []string `yaml:"fragmentShaders"`
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

	return nil
}

func (product *Product) Get() (IShaderProgram, error) {
	var errors Error.ErrorCollection

	if product.shaderProgram == nil {
		tmpProgram, err := product.Constructor(product.VertexShaders, product.FragmentShaders)
		if err != nil {
			return nil, err
		}

		product.shaderProgram = tmpProgram
	}

	return product.shaderProgram, errors.Err()
}
