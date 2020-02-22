package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type ISpotLight interface {
	Light.ILight
	Light.IPositionalLight
	Light.IDirectionalLight

	Get() SpotLight
	Set(light SpotLight)

	GetInnerCone() float32
	SetInnerCone(val float32)

	GetOuterCone() float32
	SetOuterCone(val float32)
}
