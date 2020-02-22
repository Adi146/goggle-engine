package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type DirectionalLightBase struct {
	Direction Vector.Vector3 `yaml:"direction,flow"`
}

func (light *DirectionalLightBase) GetDirection() Vector.Vector3 {
	return light.Direction
}

func (light *DirectionalLightBase) SetDirection(direction Vector.Vector3) {
	light.Direction = direction
}
