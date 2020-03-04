package UniformBufferFactory

import (
	"fmt"
	"reflect"

	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

var (
	typeFactory  = map[string]reflect.Type{}
	globalConfig FactoryConfig
)

func AddType(key string, uniformType reflect.Type) {
	typeFactory[key] = uniformType
}

func Get(key string) (UniformBuffer.IUniformBuffer, error) {
	ubo, ok := globalConfig.UniformBuffers[key]
	if !ok {
		return nil, fmt.Errorf("uniformbuffer with name %s is not configured", key)
	}

	return ubo.IUniformBuffer, nil
}

func SetConfig(config FactoryConfig) {
	globalConfig = config
}
