package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type PointLight struct {
	Position Vector.Vector3

	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}

func (light *PointLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *PointLight) GetPosition() *Vector.Vector3 {
	return &light.Position
}
