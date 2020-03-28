package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type PointLight struct {
	Position  GeometryMath.Vector3 `yaml:"position,flow"`
	Linear    float32              `yaml:"linear,flow"`
	Quadratic float32              `yaml:"quadratic,flow"`
	Ambient   GeometryMath.Vector3 `yaml:"ambient,flow"`
	Diffuse   GeometryMath.Vector3 `yaml:"diffuse,flow"`
	Specular  GeometryMath.Vector3 `yaml:"specular,flow"`
}
