package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type ISpotLight interface {
	Light.ILight
	Light.IPositionalLight
	Light.IDirectionalLight

	GetInnerCone() float32
	SetInnerCone(val float32)

	GetOuterCone() float32
	SetOuterCone(val float32)
}
