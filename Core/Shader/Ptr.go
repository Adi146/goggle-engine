package Shader

import (
	"gopkg.in/yaml.v3"
)

type Ptr struct {
	IShaderProgram
}

func (config *Ptr) UnmarshalYAML(value *yaml.Node) error {
	var name string
	value.Decode(&name)

	ubo, err := Factory.Get(name)
	if err != nil {
		return err
	}

	config.IShaderProgram = ubo
	return nil
}
