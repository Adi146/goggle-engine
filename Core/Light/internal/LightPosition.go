package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type LightPosition struct {
	Position GeometryMath.Vector3 `yaml:"position,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}

func (light *LightPosition) GetPosition() GeometryMath.Vector3 {
	return light.Position
}

func (light *LightPosition) SetPosition(pos GeometryMath.Vector3) {
	light.Position = pos
}

func (light *LightPosition) GetLinear() float32 {
	return light.Linear
}

func (light *LightPosition) SetLinear(val float32) {
	light.Linear = val
}

func (light *LightPosition) GetQuadratic() float32 {
	return light.Quadratic
}

func (light *LightPosition) SetQuadratic(val float32) {
	light.Quadratic = val
}
