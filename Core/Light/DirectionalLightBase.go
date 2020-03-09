package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type DirectionalLightBase struct {
	Direction GeometryMath.Vector3 `yaml:"direction,flow"`
}

func (light *DirectionalLightBase) GetDirection() GeometryMath.Vector3 {
	return light.Direction
}

func (light *DirectionalLightBase) SetDirection(direction GeometryMath.Vector3) {
	light.Direction = direction
}
