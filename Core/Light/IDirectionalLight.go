package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IDirectionalLight interface {
	GetDirection() GeometryMath.Vector3
	SetDirection(val GeometryMath.Vector3)
}
