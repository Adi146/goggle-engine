package UniformBufferSection

import (
	"gopkg.in/yaml.v3"
)

type Float struct {
	section
}

func (section *Float) Set(float float32) {
	section.section.set(&float)
}

func (section *Float) Get() float32 {
	val := section.section.get()
	if val == nil {
		return 0
	} else {
		return *val.(*float32)
	}
}

func (section *Float) UnmarshalYAML(value *yaml.Node) error {
	section.Set(0)
	return section.section.UnmarshalYAML(value)
}
