package DirectionalLight

import (
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

func (buff *UniformBuffer) UnmarshalYAML(value *yaml.Node) error {
	ubo := ubo.UniformBufferBase{
		Size: ubo_size,
		Type: UBO_type,
	}

	if err := value.Decode(&ubo); err != nil {
		return err
	}

	buff.UniformBufferBase = ubo
	return nil
}
