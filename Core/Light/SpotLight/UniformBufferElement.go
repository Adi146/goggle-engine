package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	element_size = 112

	element_offset = 16

	position_offset  = 0
	direction_offset = 16
	innerCone_offset = 28
	outerCone_offset = 32
	ambient_offset   = 48
	diffuse_offset   = 64
	specular_offset  = 80
	linear_offset    = 92
	quadratic_offset = 96
)

type UniformBufferElement struct {
	SpotLight `yaml:",inline"`
	ubo       *UniformBuffer

	ElementIndex int
}

func (elem *UniformBufferElement) Set(light SpotLight) {
	elem.SpotLight = light
	elem.ForceUpdate()
}

func (elem *UniformBufferElement) SetPosition(pos GeometryMath.Vector3) {
	elem.SpotLight.SetPosition(pos)
	elem.ubo.UpdateData(&pos[0], elem.getOffset()+position_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetDirection(val GeometryMath.Vector3) {
	elem.SpotLight.SetDirection(val)
	elem.ubo.UpdateData(&val[0], elem.getOffset()+direction_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetInnerCone(val float32) {
	elem.SpotLight.SetInnerCone(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+innerCone_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) SetOuterCone(val float32) {
	elem.SpotLight.SetOuterCone(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+outerCone_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) SetAmbient(color GeometryMath.Vector3) {
	elem.SpotLight.SetAmbient(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+ambient_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetDiffuse(color GeometryMath.Vector3) {
	elem.SpotLight.SetDiffuse(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+diffuse_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetSpecular(color GeometryMath.Vector3) {
	elem.SpotLight.SetSpecular(color)
	elem.ubo.UpdateData(&color[0], elem.getOffset()+specular_offset, ubo.Std140_size_vec3)
}

func (elem *UniformBufferElement) SetLinear(val float32) {
	elem.SpotLight.SetLinear(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+linear_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) SetQuadratic(val float32) {
	elem.SpotLight.SetQuadratic(val)
	elem.ubo.UpdateData(&val, elem.getOffset()+quadratic_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) ForceUpdate() {
	pos := elem.Position
	direction := elem.Direction
	innerCone := elem.InnerCone
	outerCone := elem.OuterCone
	ambient := elem.Ambient
	diffuse := elem.Diffuse
	specular := elem.Specular
	linear := elem.Linear
	quadratic := elem.Quadratic

	elem.ubo.UpdateData(&pos[0], elem.getOffset()+position_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&direction[0], elem.getOffset()+direction_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&innerCone, elem.getOffset()+innerCone_offset, ubo.Std140_size_single)
	elem.ubo.UpdateData(&outerCone, elem.getOffset()+outerCone_offset, ubo.Std140_size_single)
	elem.ubo.UpdateData(&ambient[0], elem.getOffset()+ambient_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&diffuse[0], elem.getOffset()+diffuse_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&specular[0], elem.getOffset()+specular_offset, ubo.Std140_size_vec3)
	elem.ubo.UpdateData(&linear, elem.getOffset()+linear_offset, ubo.Std140_size_single)
	elem.ubo.UpdateData(&quadratic, elem.getOffset()+quadratic_offset, ubo.Std140_size_single)
}

func (elem *UniformBufferElement) getOffset() int {
	return element_offset + elem.ElementIndex*element_size
}
