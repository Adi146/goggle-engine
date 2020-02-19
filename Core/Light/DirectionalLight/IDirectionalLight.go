package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IDirectionalLight interface {
	Light.ILight

	Set(light DirectionalLight)
	Get() DirectionalLight

	GetDirection() Vector.Vector3
	SetDirection(direction Vector.Vector3)
}
