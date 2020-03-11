package UniformBuffer

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Binding uint32 `yaml:"binding"`
	Size    int    `yaml:"size"`
	Type    Type   `yaml:"type"`

	Shaders []ShaderFactory.Config `shaders`
}

func (buff *UniformBufferBase) UnmarshalYAML(value *yaml.Node) error {
	var tmpConfig YamlConfig
	if err := value.Decode(&tmpConfig); err != nil {
		return err
	}

	tmpBuff, err := tmpConfig.Create()
	if err != nil {
		return err
	}

	*buff = tmpBuff
	return nil
}

func (yaml *YamlConfig) Create() (UniformBufferBase, error) {
	buff := NewUniformBufferBase(yaml.Size, yaml.Binding, yaml.Type)

	for _, shader := range yaml.Shaders {
		if err := shader.BindObject(&buff); err != nil {
			return buff, err
		}
	}

	return buff, nil
}
