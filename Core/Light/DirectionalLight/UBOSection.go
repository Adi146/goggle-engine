package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_direction = 0
	offset_ambient   = 16
	offset_diffuse   = 32
	offset_specular  = 48
	offset_camera    = 64
)

type UBOSection struct {
	DirectionalLight
	UniformBuffer UniformBuffer.IUniformBuffer
}

func (section *UBOSection) SetDirection(direction GeometryMath.Vector3) {
	section.DirectionalLight.SetDirection(direction)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&direction[0], offset_direction, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetAmbient(color GeometryMath.Vector3) {
	section.DirectionalLight.SetAmbient(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], offset_ambient, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetDiffuse(color GeometryMath.Vector3) {
	section.DirectionalLight.SetDiffuse(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], offset_diffuse, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) SetSpecular(color GeometryMath.Vector3) {
	section.DirectionalLight.SetSpecular(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], offset_specular, UniformBuffer.Std140_size_vec3)
	}
}

func (section *UBOSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		direction := section.Direction
		ambient := section.Ambient
		diffuse := section.Diffuse
		specular := section.Specular

		section.UniformBuffer.UpdateData(&direction[0], offset_direction, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&ambient[0], offset_ambient, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&diffuse[0], offset_diffuse, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&specular[0], offset_specular, UniformBuffer.Std140_size_vec3)
	}
}
