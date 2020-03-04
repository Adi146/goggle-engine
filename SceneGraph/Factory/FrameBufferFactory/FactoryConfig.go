package FrameBufferFactory

type FactoryConfig struct {
	FrameBuffers map[string]Product `yaml:"frameBuffers"`
}
