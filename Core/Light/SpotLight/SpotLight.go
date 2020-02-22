package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type SpotLight struct {
	Light.LightBase           `yaml:",inline"`
	Light.PositionalLightBase `yaml:",inline"`

	Direction Vector.Vector3 `yaml:"direction,flow"`

	InnerCone float32 `yaml:"innerCone,flow"`
	OuterCone float32 `yaml:"outerCone,flow"`
}

func (light *SpotLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *SpotLight) Get() SpotLight {
	return *light
}

func (light *SpotLight) Set(val SpotLight) {
	*light = val
}

func (light *SpotLight) GetDirection() Vector.Vector3 {
	return light.Direction
}

func (light *SpotLight) SetDirection(val Vector.Vector3) {
	light.Direction = val
}

func (light *SpotLight) GetInnerCone() float32 {
	return light.InnerCone
}

func (light *SpotLight) SetInnerCone(val float32) {
	light.InnerCone = val
}

func (light *SpotLight) GetOuterCone() float32 {
	return light.OuterCone
}

func (light *SpotLight) SetOuterCone(val float32) {
	light.OuterCone = val
}
