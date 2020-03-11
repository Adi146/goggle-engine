package YamlFactory

import (
	"github.com/Adi146/goggle-engine/Core/ProcessingPipeline"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/SceneFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"io/ioutil"
	"os"

	"github.com/Adi146/goggle-engine/Core/Utils/Log"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"gopkg.in/yaml.v3"
)

type config struct {
	config2 `yaml:",inline"`
}

type config2 struct {
	OpenGlLogging bool `yaml:"openGlLogging"`

	ProcessingPipelineConfig []struct {
		FrameBuffer    string               `yaml:"frameBuffer"`
		Scene          string               `yaml:"scene"`
		EnforcedShader ShaderFactory.Config `yaml:"enforcedShader"`
	} `yaml:"processingPipeline"`
}

func (config *config) UnmarshalYAML(value *yaml.Node) error {
	var frameBufferFactory FrameBufferFactory.FactoryConfig
	if err := value.Decode(&frameBufferFactory); err != nil {
		return err
	}
	FrameBufferFactory.SetConfig(frameBufferFactory)

	var shaderFactory ShaderFactory.FactoryConfig
	if err := value.Decode(&shaderFactory); err != nil {
		return err
	}
	ShaderFactory.SetConfig(shaderFactory)

	var sceneFactory SceneFactory.FactoryConfig
	if err := value.Decode(&sceneFactory); err != nil {
		return err
	}
	SceneFactory.SetConfig(sceneFactory)

	if err := value.Decode(&config.config2); err != nil {
		return err
	}

	return nil
}

func ReadConfig(file *os.File) (*Factory.Config, error) {
	var config config

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	pipelineSteps, err := config.UnmarshalProcessingPipeline()
	if err != nil {
		return nil, err
	}
	window, err := FrameBufferFactory.Get("default")
	if err != nil {
		return nil, err
	}

	if config.OpenGlLogging {
		Log.EnableDebugLogging()
	}

	return &Factory.Config{
		Pipeline: ProcessingPipeline.ProcessingPipeline{
			Steps:  pipelineSteps,
			Scenes: SceneFactory.GetAll(),
			Window: window.(Window.IWindow),
		},
	}, nil
}

func (config *config) UnmarshalProcessingPipeline() ([]ProcessingPipeline.Step, error) {
	var Pipeline []ProcessingPipeline.Step

	for _, stepConfig := range config.ProcessingPipelineConfig {
		frameBuffer, err := FrameBufferFactory.Get(stepConfig.FrameBuffer)
		if err != nil {
			return nil, err
		}
		scene, err := SceneFactory.Get(stepConfig.Scene)
		if err != nil {
			return nil, err
		}

		Pipeline = append(
			Pipeline,
			ProcessingPipeline.Step{
				Scene:          scene,
				FrameBuffer:    frameBuffer,
				EnforcedShader: stepConfig.EnforcedShader.IShaderProgram,
			},
		)
	}

	return Pipeline, nil
}
