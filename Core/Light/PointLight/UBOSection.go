package PointLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	section_size = 80

	offset_position  = 0
	offset_ambient   = 16
	offset_diffuse   = 32
	offset_specular  = 48
	offset_linear    = 60
	offset_quadratic = 64
)

type UBOSection struct {
	PointLight
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *UBOSection) Set(light PointLight) {
	section.PointLight = light
	section.ForceUpdate()
}

func (section *UBOSection) SetPosition(pos GeometryMath.Vector3) {
	section.PointLight.SetPosition(pos)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetAmbient(color GeometryMath.Vector3) {
	section.PointLight.SetAmbient(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_ambient, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetDiffuse(color GeometryMath.Vector3) {
	section.PointLight.SetDiffuse(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_diffuse, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetSpecular(color GeometryMath.Vector3) {
	section.PointLight.SetSpecular(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_specular, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetLinear(val float32) {
	section.PointLight.SetLinear(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_linear, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) SetQuadratic(val float32) {
	section.PointLight.SetQuadratic(val)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&val, section.Offset+offset_quadratic, UniformBuffer.Std140_size_single)
	}
}

func (section *UBOSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		pos := section.Position
		ambient := section.Ambient
		diffuse := section.Diffuse
		specular := section.Specular
		linear := section.Linear
		quadratic := section.Quadratic

		section.UniformBuffer.UpdateData(&pos[0], section.Offset+offset_position, UniformBuffer.Std140_size_vec3)
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
