package UniformBufferFactory

import (
	"fmt"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

var (
	typeFactory  = map[string]reflect.Type{}
	globalConfig UniformBuffersConfig
)

func AddType(key string, uniformType reflect.Type) {
	typeFactory[key] = uniformType
}

func Get(key string) (UniformBuffer.IUniformBuffer, error) {
	return globalConfig.Get(key)
}

func SetConfig(config UniformBuffersConfig) {
	globalConfig = config
}

type UniformBuffersConfig struct {
	UniformBufferConfig   map[string]UniformBufferConfig `yaml:"uniformBuffers"`
	DecodedUniformBuffers map[string]UniformBuffer.IUniformBuffer
}

func (config *UniformBuffersConfig) Get(name string) (UniformBuffer.IUniformBuffer, error) {
	if uniformBuffer, ok := config.DecodedUniformBuffers[name]; ok {
		return uniformBuffer, nil
	}

	uniformBufferConfig, ok := config.UniformBufferConfig[name]
	if !ok {
		return nil, fmt.Errorf("uniformbuffer %s is not configured", name)
	}

	uniformBuffer, err := uniformBufferConfig.Unmarshal()
	if err == nil {
		config.DecodedUniformBuffers[name] = uniformBuffer
	}

	return uniformBuffer, err
}

type UniformBufferConfig struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

func (config *UniformBufferConfig) Unmarshal() (UniformBuffer.IUniformBuffer, error) {
	uniformBufferType, ok := typeFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("uniformbuffer type %s is not in factory", config.Type)
	}

	uniformBuffer := reflect.New(uniformBufferType).Interface().(UniformBuffer.IUniformBuffer)

	config.Config.Decode(uniformBuffer)

	if err := uniformBuffer.Init(); err != nil {
		return nil, err
	}

	return uniformBuffer, nil
}
