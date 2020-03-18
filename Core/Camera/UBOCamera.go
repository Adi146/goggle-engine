package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	ubo_size                    = 2 * UniformBuffer.Std140_size_mat4
	UBO_type UniformBuffer.Type = "camera"
)

type UBOCamera CameraSection

func (camera *UBOCamera) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	((*CameraSection)(camera)).SetProjectionMatrix(matrix)
}

func (camera *UBOCamera) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	((*CameraSection)(camera)).SetViewMatrix(matrix)
}
