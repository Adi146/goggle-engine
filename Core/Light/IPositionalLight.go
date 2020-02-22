package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type IPositionalLight interface {
	GetPosition() Vector.Vector3
	SetPosition(pos Vector.Vector3)

	GetLinear() float32
	SetLinear(val float32)

	GetQuadratic() float32
	SetQuadratic(val float32)
}
