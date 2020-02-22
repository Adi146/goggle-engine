package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IPointLight interface {
	Light.ILight
	Light.IPositionalLight

	Get() PointLight
	Set(light PointLight)
}
