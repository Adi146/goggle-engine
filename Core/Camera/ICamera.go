package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type ICamera interface {
	Get() Camera
	Set(camera Camera)

	GetViewMatrix() GeometryMath.Matrix4x4
	SetViewMatrix(matrix GeometryMath.Matrix4x4)

	GetProjectionMatrix() GeometryMath.Matrix4x4
	SetProjectionMatrix(matrix GeometryMath.Matrix4x4)
}
