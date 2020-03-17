package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type SpotLight struct {
	Light.LightBase            `yaml:",inline"`
	Light.PositionalLightBase  `yaml:",inline"`
	Light.DirectionalLightBase `yaml:",inline"`

	InnerCone float32 `yaml:"innerCone"`
	OuterCone float32 `yaml:"outerCone"`
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
