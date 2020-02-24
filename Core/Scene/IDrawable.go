package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IDrawable interface {
	Draw() error
	GetPosition() *Vector.Vector3
}
