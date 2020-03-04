package FrameBufferFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"gopkg.in/yaml.v3"
	"reflect"
)

type Product struct {
	FrameBuffer.IFrameBuffer
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

	frameBufferType, ok := typeFactory[tmpProduct.Type]
	if !ok {
		return fmt.Errorf("framebuffer type %s is not in factory", tmpProduct.Type)
	}

	frameBuffer := reflect.New(frameBufferType).Interface().(FrameBuffer.IFrameBuffer)

	if err := tmpProduct.Config.Decode(frameBuffer); err != nil {
		return err
	}

	if err := frameBuffer.Init(); err != nil {
		return err
	}

	product.IFrameBuffer = frameBuffer
	return nil
}
