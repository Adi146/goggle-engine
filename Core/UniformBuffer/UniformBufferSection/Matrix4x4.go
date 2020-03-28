package UniformBufferSection

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

type Matrix4x4 struct {
	section
}

func (section *Matrix4x4) Set(mat GeometryMath.Matrix4x4) {
	section.section.set(&mat)
}

func (section *Matrix4x4) Get() GeometryMath.Matrix4x4 {
	val := section.section.get()
	if val == nil {
		return GeometryMath.Matrix4x4{}
	} else {
		return *val.(*GeometryMath.Matrix4x4)
	}
}

func (section *Matrix4x4) UnmarshalYAML(value *yaml.Node) error {
	section.Set(GeometryMath.Matrix4x4{})
	return section.section.UnmarshalYAML(value)
}
