package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
)

type ICamera interface {
	Get() Camera
	Set(camera Camera)

	GetViewMatrix() Matrix.Matrix4x4
	SetViewMatrix(matrix Matrix.Matrix4x4)

	GetProjectionMatrix() Matrix.Matrix4x4
	SetProjectionMatrix(matrix Matrix.Matrix4x4)
}
