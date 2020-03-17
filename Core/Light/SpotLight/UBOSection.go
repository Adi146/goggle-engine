package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	section_size = 112

	offset_position  = 0
	offset_direction = 16
	offset_innerCone = 28
	offset_outerCone = 32
	offset_ambient   = 48
	offset_diffuse   = 64
	offset_specular  = 80
	offset_linear    = 92
	offset_quadratic = 96
)

type UBOSection struct {
	SpotLight
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *UBOSection) SetPosition(pos GeometryMath.Vector3) {
	section.SpotLight.SetPosition(pos)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetDirection(val GeometryMath.Vector3) {
	section.SpotLight.SetDirection(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val[0], section.Offset+offset_direction, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetInnerCone(val float32) {
	section.SpotLight.SetInnerCone(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_innerCone, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) SetOuterCone(val float32) {
	section.SpotLight.SetOuterCone(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_outerCone, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) SetAmbient(color GeometryMath.Vector3) {
	section.SpotLight.SetAmbient(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_ambient, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetDiffuse(color GeometryMath.Vector3) {
	section.SpotLight.SetDiffuse(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_diffuse, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetSpecular(color GeometryMath.Vector3) {
	section.SpotLight.SetSpecular(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_specular, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetLinear(val float32) {
	section.SpotLight.SetLinear(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_linear, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) SetQuadratic(val float32) {
	section.SpotLight.SetQuadratic(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_quadratic, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		pos := section.Position
		direction := section.Direction
		innerCone := section.InnerCone
		outerCone := section.OuterCone
		ambient := section.Ambient
		diffuse := section.Diffuse
		specular := section.Specular
		linear := section.Linear
		quadratic := section.Quadratic

		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&direction[0], section.Offset+offset_direction, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&innerCone, section.Offset+offset_innerCone, UniformBuffer.Std140_size_single)
		section.UniformBuffer.UpdateData(&outerCone, section.Offset+offset_outerCone, UniformBuffer.Std140_size_single)
		section.UniformBuffer.UpdateData(&ambient[0], section.Offset+offset_ambient, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&diffuse[0], section.Offset+offset_diffuse, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&specular[0], section.Offset+offset_specular, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&linear, section.Offset+offset_linear, UniformBuffer.Std140_size_single)
		section.UniformBuffer.UpdateData(&quadratic, section.Offset+offset_quadratic, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *UBOSection) GetSize() int {
	return section_size
}
