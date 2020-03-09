package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type PositionalLightBase struct {
	Position GeometryMath.Vector3 `yaml:"position,flow"`

	Linear    float32 `yaml:"linear,flow"`
	Quadratic float32 `yaml:"quadratic,flow"`
}

func (light *PositionalLightBase) GetPosition() GeometryMath.Vector3 {
	return light.Position
}

func (light *PositionalLightBase) SetPosition(pos GeometryMath.Vector3) {
	light.Position = pos
}

func (light *PositionalLightBase) GetLinear() float32 {
	return light.Linear
}

func (light *PositionalLightBase) SetLinear(val float32) {
	light.Linear = val
}

func (light *PositionalLightBase) GetQuadratic() float32 {
	return light.Quadratic
}

func (light *PositionalLightBase) SetQuadratic(val float32) {
	light.Quadratic = val
}
