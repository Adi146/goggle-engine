package UniformBufferFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
	"reflect"
)

type Product struct {
	UniformBuffer.IUniformBuffer
}

type tmpProduct struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var tmpProduct tmpProduct
	if err := value.Decode(&tmpProduct); err != nil {
		return err
	}

	uniformBufferType, ok := typeFactory[tmpProduct.Type]
	if !ok {
		return fmt.Errorf("uniformbuffer type %s is not in factory", tmpProduct.Type)
	}

	uniformBuffer := reflect.New(uniformBufferType).Interface().(UniformBuffer.IUniformBuffer)

	if err := tmpProduct.Config.Decode(uniformBuffer); err != nil {
		return err
	}

	if err := uniformBuffer.Init(); err != nil {
		return err
	}

	product.IUniformBuffer = uniformBuffer
	return nil
}
