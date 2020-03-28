package UniformBufferSection

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

type Vector3 struct {
	section
}

func (section *Vector3) Set(vec GeometryMath.Vector3) {
	section.section.set(&vec)
}

func (section *Vector3) Get() GeometryMath.Vector3 {
	val := section.section.get()
	if val == nil {
		return GeometryMath.Vector3{}
	} else {
		return *val.(*GeometryMath.Vector3)
	}
}

func (section *Vector3) UnmarshalYAML(value *yaml.Node) error {
	section.Set(GeometryMath.Vector3{})
	return section.section.UnmarshalYAML(value)
}
