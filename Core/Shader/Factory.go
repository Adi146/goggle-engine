package Shader

import (
	"fmt"
)

var Factory = factory{
	Products:     make(map[string]*Product),
	Constructors: make(map[string]func([]string, []string) (IShaderProgram, error)),
}

type factory struct {
	Products     map[string]*Product `yaml:"shaders"`
	Constructors map[string]func([]string, []string) (IShaderProgram, error)
}

func (factory *factory) AddConstructor(key string, constructor func([]string, []string) (IShaderProgram, error)) {
	factory.Constructors[key] = constructor
}

func (factory *factory) Get(key string) (IShaderProgram, error) {
	product, ok := factory.Products[key]
	if !ok {
		return nil, fmt.Errorf("shader with name %s is not configured", key)
	}

	return product.Get()
}
