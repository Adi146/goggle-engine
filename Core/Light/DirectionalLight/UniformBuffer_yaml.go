package DirectionalLight

import (
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

func (buff *UniformBuffer) UnmarshalYAML(value *yaml.Node) error {
	tmpConfig := ubo.YamlConfig{
		Size: ubo_size,
		Type: UBO_type,
	}
	if err := value.Decode(&tmpConfig); err != nil {
		return err
	}

	tmpBuff, err := tmpConfig.Create()
	if err != nil {
		return err
	}

	buff.UniformBufferBase = tmpBuff
	return nil
}
