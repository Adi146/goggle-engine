package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_direction        = 0
	offset_ambient          = 16
	offset_diffuse          = 32
	offset_specular         = 48
	offset_projectionMatrix = 64
	offset_viewMatrix       = 128

	ubo_size          = 192
	UBO_type ubo.Type = "directionalLight"
)

type UniformBuffer struct {
	DirectionalLight
	ubo.UniformBufferBase
}

func (buff *UniformBuffer) Set(light DirectionalLight) {
	buff.DirectionalLight = light
	buff.ForceUpdate()
}

func (buff *UniformBuffer) SetDirection(direction GeometryMath.Vector3) {
	buff.DirectionalLight.SetDiffuse(direction)
	buff.UpdateData(&direction[0], offset_direction, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetAmbient(color GeometryMath.Vector3) {
	buff.DirectionalLight.SetAmbient(color)
	buff.UpdateData(&color[0], offset_ambient, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetDiffuse(color GeometryMath.Vector3) {
	buff.DirectionalLight.SetDiffuse(color)
	buff.UpdateData(&color[0], offset_diffuse, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetSpecular(color GeometryMath.Vector3) {
	buff.DirectionalLight.SetSpecular(color)
	buff.UpdateData(&color[0], offset_specular, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetProjectionMatrix(matrix GeometryMath.Matrix4x4) {
	buff.DirectionalLight.SetProjectionMatrix(matrix)
	buff.UpdateData(&matrix[0][0], offset_projectionMatrix, ubo.Std140_size_mat4)
}

func (buff *UniformBuffer) SetViewMatrix(matrix GeometryMath.Matrix4x4) {
	buff.DirectionalLight.SetViewMatrix(matrix)
	buff.UpdateData(&matrix[0][0], offset_viewMatrix, ubo.Std140_size_mat4)
}

func (buff *UniformBuffer) ForceUpdate() {
	direction := buff.Direction
	ambient := buff.Ambient
	diffuse := buff.Diffuse
	specular := buff.Specular
	projectionMatrix := buff.ProjectionMatrix
	viewMatrix := buff.ViewMatrix

	buff.UpdateData(&direction[0], offset_direction, ubo.Std140_size_vec3)
	buff.UpdateData(&ambient[0], offset_ambient, ubo.Std140_size_vec3)
	buff.UpdateData(&diffuse[0], offset_diffuse, ubo.Std140_size_vec3)
	buff.UpdateData(&specular[0], offset_specular, ubo.Std140_size_vec3)
	buff.UpdateData(&projectionMatrix[0][0], offset_projectionMatrix, ubo.Std140_size_mat4)
	buff.UpdateData(&viewMatrix[0][0], offset_viewMatrix, ubo.Std140_size_mat4)
}
