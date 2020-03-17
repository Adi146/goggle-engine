package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type PointLight struct {
	Light.LightBase           `yaml:",inline"`
	Light.PositionalLightBase `yaml:",inline"`
}
