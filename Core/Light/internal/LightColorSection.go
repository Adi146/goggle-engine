package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	offset_ambient  = 0
	offset_diffuse  = 16
	offset_specular = 32

	Size_lightColorSection = 48
)

type LightColorSection struct {
	LightColor
	UniformBuffer UniformBuffer.IUniformBuffer
	Offset        int
}

func (section *LightColorSection) SetAmbient(color GeometryMath.Vector3) {
	section.LightColor.SetAmbient(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_ambient, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightColorSection) SetDiffuse(color GeometryMath.Vector3) {
	section.LightColor.SetDiffuse(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_diffuse, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightColorSection) SetSpecular(color GeometryMath.Vector3) {
	section.LightColor.SetSpecular(color)
	if section.UniformBuffer != nil {
		section.UniformBuffer.UpdateData(&color[0], section.Offset+offset_specular, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightColorSection) ForceUpdate() {
	if section.UniformBuffer != nil {
		ambient := section.Ambient
		diffuse := section.Diffuse
		specular := section.Specular

		section.UniformBuffer.UpdateData(&ambient[0], section.Offset+offset_ambient, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&diffuse[0], section.Offset+offset_diffuse, UniformBuffer.Std140_size_vec3)
		section.UniformBuffer.UpdateData(&specular[0], section.Offset+offset_specular, UniformBuffer.Std140_size_vec3)
	}
}

func (section *LightColorSection) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	section.UniformBuffer = ubo
	section.Offset = offset
}

func (section *LightColorSection) GetSize() int {
	return Size_lightColorSection
}

func (section *LightColorSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightColor); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
