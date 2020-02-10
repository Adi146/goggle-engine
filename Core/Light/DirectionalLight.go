package Light

import (
	"math"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type DirectionalLight struct {
	Direction Vector.Vector3

	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`
}

func (light *DirectionalLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *DirectionalLight) GetPosition() *Vector.Vector3 {
	return light.Direction.MulScalar(-math.MaxFloat32)
}
