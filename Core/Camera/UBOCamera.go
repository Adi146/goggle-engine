package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	ubo_size                    = 2 * UniformBuffer.Std140_size_mat4
	UBO_type UniformBuffer.Type = "camera"
)

type UBOCamera UBOSection

func (camera *UBOCamera) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	((*UBOSection)(camera)).SetProjectionMatrix(matrix)
}

func (camera *UBOCamera) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	((*UBOSection)(camera)).SetViewMatrix(matrix)
}
