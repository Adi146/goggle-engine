package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type PointLight struct {
	Position Vector.Vector3

	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}
