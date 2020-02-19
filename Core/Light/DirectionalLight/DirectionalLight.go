package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
	"math"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type DirectionalLight struct {
	Light.LightBase `yaml:",inline"`

	Direction Vector.Vector3 `yaml:"direction,flow"`
}

func (light *DirectionalLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *DirectionalLight) GetPosition() *Vector.Vector3 {
	return light.Direction.MulScalar(-math.MaxFloat32)
}

func (light *DirectionalLight) Set(val DirectionalLight) {
	*light = val
}

func (light *DirectionalLight) Get() DirectionalLight {
	return *light
}

func (light *DirectionalLight) GetDirection() Vector.Vector3 {
	return light.Direction
}

func (light *DirectionalLight) SetDirection(direction Vector.Vector3) {
	light.Direction = direction
}
