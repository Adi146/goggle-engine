package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type DirectionalLight struct {
	Direction GeometryMath.Vector3 `yaml:"direction,flow"`
	Ambient   GeometryMath.Vector3 `yaml:"ambient,flow"`
	Diffuse   GeometryMath.Vector3 `yaml:"diffuse,flow"`
	Specular  GeometryMath.Vector3 `yaml:"specular,flow"`
}
