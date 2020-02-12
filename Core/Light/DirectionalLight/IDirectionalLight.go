package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IDirectionalLight interface {
	Set(light DirectionalLight)
	Get() DirectionalLight

	GetDirection() Vector.Vector3
	SetDirection(direction Vector.Vector3)

	GetAmbient() Vector.Vector3
	SetAmbient(color Vector.Vector3)

	GetDiffuse() Vector.Vector3
	SetDiffuse(color Vector.Vector3)

	GetSpecular() Vector.Vector3
	SetSpecular(color Vector.Vector3)
}
