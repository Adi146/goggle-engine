package FrameBufferFactory

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"gopkg.in/yaml.v3"
)

type Config struct {
	FrameBuffer.IFrameBuffer
}

func (config *Config) UnmarshalYAML(value *yaml.Node) error {
	var name string
	value.Decode(&name)

	fbo, err := Get(name)
	if err != nil {
		return err
	}

	config.IFrameBuffer = fbo
	return nil
}
