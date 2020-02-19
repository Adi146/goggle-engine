package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IPointLight interface {
	Light.ILight

	Get() PointLight
	Set(light PointLight)

	GetPosition() Vector.Vector3
	SetPosition(pos Vector.Vector3)

	GetLinear() float32
	SetLinear(val float32)

	GetQuadratic() float32
	SetQuadratic(val float32)
}
