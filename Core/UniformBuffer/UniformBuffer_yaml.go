package UniformBuffer

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"gopkg.in/yaml.v3"
)

func (buff *UniformBuffer) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Binding uint32 `yaml:"binding"`
		Size    int    `yaml:"size"`
		Type    Type   `yaml:"type"`

		Shaders []ShaderFactory.Config `shaders`
	}{
		Binding: buff.Binding,
		Size:    buff.Size,
		Type:    buff.Type,
	}

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpBuff := NewUniformBufferBase(yamlConfig.Size, yamlConfig.Binding, yamlConfig.Type)
	for _, shader := range yamlConfig.Shaders {
		if err := shader.BindObject(&tmpBuff); err != nil {
			return err
		}
	}

	*buff = tmpBuff
	return nil
}
