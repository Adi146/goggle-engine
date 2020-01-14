package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type SpotLight struct {
	Position  Vector.Vector3
	Direction Vector.Vector3

	InnerCone float32 `yaml:"innerCone,flow"`
	OuterCone float32 `yaml:"outerCone,flow"`

	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}
