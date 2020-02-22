package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type IDirectionalLight interface {
	GetDirection() Vector.Vector3
	SetDirection(val Vector.Vector3)
}
