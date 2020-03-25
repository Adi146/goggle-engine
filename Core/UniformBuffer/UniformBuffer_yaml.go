package UniformBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

func (buff *UniformBuffer) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Binding uint32 `yaml:"binding"`
		Size    int    `yaml:"size"`
		Type    Type   `yaml:"type"`

		Shaders []Shader.Ptr `shaders`
	}{
		Binding: buff.Binding,
		Size:    buff.Size,
		Type:    buff.Type,
	}

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpBuff, err := NewUniformBufferBase(yamlConfig.Size, yamlConfig.Binding, yamlConfig.Type)
	if err != nil {
		return err
	}
	for _, shader := range yamlConfig.Shaders {
		if err := shader.BindObject(&tmpBuff); err != nil {
			return err
		}
	}

	*buff = tmpBuff
	return nil
}
