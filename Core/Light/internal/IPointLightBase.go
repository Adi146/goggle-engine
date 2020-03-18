package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type ILightPosition interface {
	GetPosition() GeometryMath.Vector3
	SetPosition(pos GeometryMath.Vector3)

	GetLinear() float32
	SetLinear(val float32)

	GetQuadratic() float32
	SetQuadratic(val float32)
}
