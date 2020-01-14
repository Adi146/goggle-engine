package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
)

type ICamera interface {
	GetViewMatrix() *Matrix.Matrix4x4
	GetProjectionMatrix() *Matrix.Matrix4x4
	Tick(timeDelta float32)
}
