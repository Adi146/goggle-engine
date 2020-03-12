package FrameBufferFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Window"
	"reflect"
)

var (
	typeFactory = map[string]reflect.Type{
		"sdlWindow": reflect.TypeOf((*Window.SDLWindow)(nil)).Elem(),
	}
	globalConfig FactoryConfig
)

func AddType(key string, frameBufferType reflect.Type) {
	typeFactory[key] = frameBufferType
}

func Get(key string) (FrameBuffer.IFrameBuffer, error) {
	fbo, ok := globalConfig.FrameBuffers[key]
	if !ok {
		return nil, fmt.Errorf("framebuffer with name %s is not configured", key)
	}

	return fbo.IFrameBuffer, nil
}

func SetConfig(config FactoryConfig) {
	globalConfig = config
}
