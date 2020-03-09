package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IDirectionalLight interface {
	Light.ILight
	Light.IDirectionalLight

	GetViewMatrix() GeometryMath.Matrix4x4
	SetViewMatrix(matrix GeometryMath.Matrix4x4)

	GetProjectionMatrix() GeometryMath.Matrix4x4
	SetProjectionMatrix(matrix GeometryMath.Matrix4x4)

	Set(light DirectionalLight)
	Get() DirectionalLight
}
