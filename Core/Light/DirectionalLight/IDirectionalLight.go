package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IDirectionalLight interface {
	Light.ILight
	Light.IDirectionalLight

	GetViewMatrix() Matrix.Matrix4x4
	SetViewMatrix(matrix Matrix.Matrix4x4)

	GetProjectionMatrix() Matrix.Matrix4x4
	SetProjectionMatrix(matrix Matrix.Matrix4x4)

	Set(light DirectionalLight)
	Get() DirectionalLight
}
