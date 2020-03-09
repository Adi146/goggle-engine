package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	element_size = 80

	element_offset = 16

	position_offset  = 0
	ambient_offset   = 16
	diffuse_offset   = 32
	specular_offset  = 48
	linear_offset    = 60
	quadratic_offset = 64
)

type UniformBufferElement struct {
	PointLight `yaml:",inline"`
	ubo        *UniformBuffer

	ElementIndex int
}

func (elem *UniformBufferElement) Set(light PointLight) {
	elem.PointLight = light
	elem.ForceUpdate()
}

func (elem *UniformBufferElement) SetPosition(pos GeometryMath.Vector3) {
	elem.PointLight.SetPosition(pos)
	elem.ubo.UpdateData(&pos[0], elem.getOffset()+position_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetAmbient(color GeometryMath.Vector3) {
	elem.PointLight.SetAmbient(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+ambient_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetDiffuse(color GeometryMath.Vector3) {
	elem.PointLight.SetDiffuse(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+diffuse_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetSpecular(color GeometryMath.Vector3) {
	elem.PointLight.SetSpecular(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+specular_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetLinear(val float32) {
	elem.PointLight.SetLinear(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+linear_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) SetQuadratic(val float32) {
	elem.PointLight.SetQuadratic(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+quadratic_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) ForceUpdate() {
	pos := elem.Position
	ambient := elem.Ambient
	diffuse := elem.Diffuse
	specular := elem.Specular
	linear := elem.Linear
	quadratic := elem.Quadratic

	elem.ubo.UpdateData(&pos[0], elem.getOffset()+position_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&ambient[0], elem.getOffset()+ambient_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&diffuse[0], elem.getOffset()+diffuse_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&specular[0], elem.getOffset()+specular_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&linear, elem.getOffset()+linear_offset, ubo.Std140_size_single)
	elem.ubo.UpdateData(&quadratic, elem.getOffset()+quadratic_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) getOffset() int {
	return element_offset + elem.ElementIndex*element_size
}
