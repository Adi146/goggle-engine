package YamlFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Window"
	"gopkg.in/yaml.v3"
	"reflect"
)

var FrameBufferFactory = map[string]reflect.Type{
	"sdlWindow": reflect.TypeOf((*Window.SDLWindow)(nil)).Elem(),
}

type FrameBuffersConfig struct {
	FrameBufferConfig   map[string]FrameBufferConfig `yaml:"frameBuffers"`
	DecodedFrameBuffers map[string]FrameBuffer.IFrameBuffer
}

func (config *FrameBuffersConfig) Get(name string) (FrameBuffer.IFrameBuffer, error) {
	if frameBuffer, ok := config.DecodedFrameBuffers[name]; ok {
		return frameBuffer, nil
	}

	frameBufferConfig, ok := config.FrameBufferConfig[name]
	if !ok {
		return nil, fmt.Errorf("framebuffer %s is not configured", name)
	}

	frameBuffer, err := frameBufferConfig.Unmarshal()
	if err == nil {
		config.DecodedFrameBuffers[name] = frameBuffer
	}

	return frameBuffer, err
}

type FrameBufferConfig struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

func (config *FrameBufferConfig) Unmarshal() (FrameBuffer.IFrameBuffer, error) {
	frameBufferType, ok := FrameBufferFactory[config.Type]
	if !ok {
		return nil, fmt.Errorf("framebuffer type %s is not in factory", config.Type)
	}

	frameBuffer := reflect.New(frameBufferType).Interface().(FrameBuffer.IFrameBuffer)

	config.Config.Decode(frameBuffer)

	if err := frameBuffer.Init(); err != nil {
		return nil, err
	}

	return frameBuffer, nil
}
