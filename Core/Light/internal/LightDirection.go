package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type LightDirection struct {
	Direction GeometryMath.Vector3 `yaml:"direction,flow"`
}

func (light *LightDirection) GetDirection() GeometryMath.Vector3 {
	return light.Direction
}

func (light *LightDirection) SetDirection(direction GeometryMath.Vector3) {
	light.Direction = direction
}
