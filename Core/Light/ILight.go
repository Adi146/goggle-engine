package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type ILight interface {
	GetAmbient() Vector.Vector3
	SetAmbient(color Vector.Vector3)

	GetDiffuse() Vector.Vector3
	SetDiffuse(color Vector.Vector3)

	GetSpecular() Vector.Vector3
	SetSpecular(color Vector.Vector3)
}
