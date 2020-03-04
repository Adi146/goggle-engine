package UniformBufferFactory

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

type Config struct {
	UniformBuffer.IUniformBuffer
}

func (config *Config) UnmarshalYAML(value *yaml.Node) error {
	var name string
	value.Decode(&name)

	ubo, err := Get(name)
	if err != nil {
		return err
	}

	config.IUniformBuffer = ubo
	return nil
}
