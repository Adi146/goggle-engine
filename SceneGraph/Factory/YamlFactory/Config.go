package YamlFactory

import (
	"io/ioutil"
	"os"

	"github.com/Adi146/goggle-engine/Core/UniformBuffer"

	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/ProcessingPipeline"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Utils/Log"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/UniformBufferFactory"
	"gopkg.in/yaml.v3"
)

type config struct {
	ScenesConfig                              `yaml:",inline"`
	ShaderFactory.ShadersConfig               `yaml:",inline"`
	FrameBuffersConfig                        `yaml:",inline"`
	UniformBufferFactory.UniformBuffersConfig `yaml:",inline"`

	OpenGlLogging bool `yaml:"openGlLogging"`

	ProcessingPipelineConfig []struct {
		Shader      string `yaml:"shader"`
		FrameBuffer string `yaml:"frameBuffer"`
		Scene       string `yaml:"scene"`
	} `yaml:"processingPipeline"`
}

func ReadConfig(file *os.File) (*Factory.Config, error) {
	config := config{
		ScenesConfig: ScenesConfig{
			DecodedScenes: map[string]Scene.IScene{},
		},
		ShadersConfig: ShaderFactory.ShadersConfig{
			DecodedShaders: map[string]Shader.IShaderProgram{},
		},
		FrameBuffersConfig: FrameBuffersConfig{
			DecodedFrameBuffers: map[string]FrameBuffer.IFrameBuffer{},
		},
		UniformBuffersConfig: UniformBufferFactory.UniformBuffersConfig{
			DecodedUniformBuffers: map[string]UniformBuffer.IUniformBuffer{},
		},
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	UniformBufferFactory.SetConfig(config.UniformBuffersConfig)
	ShaderFactory.SetConfig(config.ShadersConfig)

	pipelineSteps, err := config.UnmarshalProcessingPipeline()
	if err != nil {
		return nil, err
	}
	window, err := config.FrameBuffersConfig.Get("default")
	if err != nil {
		return nil, err
	}

	if config.OpenGlLogging {
		Log.EnableDebugLogging()
	}

	return &Factory.Config{
		Pipeline: ProcessingPipeline.ProcessingPipeline{
			Steps:  pipelineSteps,
			Scenes: config.GetScenes(),
			Window: window.(Window.IWindow),
		},
	}, nil
}

func (config *config) UnmarshalProcessingPipeline() ([]ProcessingPipeline.ProcessingPipelineStep, error) {
	var Pipeline []ProcessingPipeline.ProcessingPipelineStep

	for _, stepConfig := range config.ProcessingPipelineConfig {
		frameBuffer, err := config.FrameBuffersConfig.Get(stepConfig.FrameBuffer)
		if err != nil {
			return nil, err
		}
		scene, err := config.ScenesConfig.Get(stepConfig.Scene)
		if err != nil {
			return nil, err
		}
		shader, err := config.ShadersConfig.Get(stepConfig.Shader)
		if err != nil {
			return nil, err
		}

		Pipeline = append(
			Pipeline,
			ProcessingPipeline.ProcessingPipelineStep{
				Scene:       scene,
				FrameBuffer: frameBuffer,
				Shader:      shader,
			},
		)
	}

	return Pipeline, nil
}
