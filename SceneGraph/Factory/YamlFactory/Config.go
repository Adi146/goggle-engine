package YamlFactory

import (
	"github.com/Adi146/goggle-engine/Core/ProcessingPipeline"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/SceneFactory"
	"io/ioutil"
	"os"

	"github.com/Adi146/goggle-engine/Core/Utils/Log"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"gopkg.in/yaml.v3"
)

type config struct {
	config2 `yaml:",inline"`
}

type config2 struct {
	OpenGlLogging bool `yaml:"openGlLogging"`

	ProcessingPipelineConfig []struct {
		FrameBuffer    string     `yaml:"frameBuffer"`
		Scene          string     `yaml:"scene"`
		EnforcedShader Shader.Ptr `yaml:"enforcedShader"`
	} `yaml:"processingPipeline"`
}

func (config *config) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&Shader.Factory); err != nil {
		return err
	}

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

	if config.OpenGlLogging {
		Log.EnableDebugLogging()
	}

	return &Factory.Config{
		Pipeline: ProcessingPipeline.ProcessingPipeline{
			Scene: SceneFactory.GetAll()[0],
		},
	}, nil
}
