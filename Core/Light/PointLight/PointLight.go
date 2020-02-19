package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type PointLight struct {
	Light.LightBase `yaml:",inline"`

	Position Vector.Vector3 `yaml:"position,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}

func (light *PointLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *PointLight) Get() PointLight {
	return *light
}

func (light *PointLight) Set(val PointLight) {
	*light = val
}

func (light *PointLight) GetPosition() Vector.Vector3 {
	return light.Position
}

func (light *PointLight) SetPosition(pos Vector.Vector3) {
	light.Position = pos
}

func (light *PointLight) GetLinear() float32 {
	return light.Linear
}

func (light *PointLight) SetLinear(val float32) {
	light.Linear = val
}

func (light *PointLight) GetQuadratic() float32 {
	return light.Quadratic
}

func (light *PointLight) SetQuadratic(val float32) {
	light.Quadratic = val
}
