package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type ILightDirection interface {
	GetDirection() GeometryMath.Vector3
	SetDirection(val GeometryMath.Vector3)
}
