package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type ISpotLight interface {
	Light.IPositionalLight

	Get() SpotLight
	Set(light SpotLight)

	GetDirection() Vector.Vector3
	SetDirection(val Vector.Vector3)

	GetInnerCone() float32
	SetInnerCone(val float32)

	GetOuterCone() float32
	SetOuterCone(val float32)
}
