package UniformBufferSection

import (
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/Utils"
	"gopkg.in/yaml.v3"
)

type section struct {
	value         interface{}
	uniformBuffer UniformBuffer.IUniformBuffer
	offset        int
}

func (section *section) set(val interface{}) {
	section.value = val

	section.ForceUpdate()
}

func (section *section) get() interface{} {
	return section.value
}

func (section *section) ForceUpdate() {
	if section.uniformBuffer != nil {
		section.uniformBuffer.UpdateData(section.value, section.offset)
	}
}

func (section *section) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.uniformBuffer = ubo
	section.offset = offset

	section.ForceUpdate()
}

func (section *section) GetSize() int {
	valSize := Utils.SizeOf(section.value)
	if valSize%16 == 0 {
		return valSize
	} else {
		return ((valSize / 16) + 1) * 16
	}
}

func (section *section) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(section.value); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
