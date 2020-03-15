package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	offset_direction = 0
	offset_ambient   = 16
	offset_diffuse   = 32
	offset_specular  = 48
	offset_camera    = 64

	ubo_size                    = 192
	UBO_type UniformBuffer.Type = "directionalLight"
)

type UBOSection struct {
	DirectionalLight
	*UniformBuffer.UniformBufferBase
	Camera.UBOSection
}

func (section *UBOSection) SetDirection(direction GeometryMath.Vector3) {
	section.DirectionalLight.SetDiffuse(direction)
	section.UpdateData(&direction[0], offset_direction, UniformBuffer.Std140_size_vec3)
}

func (section *UBOSection) SetAmbient(color GeometryMath.Vector3) {
	section.DirectionalLight.SetAmbient(color)
	section.UpdateData(&color[0], offset_ambient, UniformBuffer.Std140_size_vec3)
}

func (section *UBOSection) SetDiffuse(color GeometryMath.Vector3) {
	section.DirectionalLight.SetDiffuse(color)
	section.UpdateData(&color[0], offset_diffuse, UniformBuffer.Std140_size_vec3)
}

func (section *UBOSection) SetSpecular(color GeometryMath.Vector3) {
	section.DirectionalLight.SetSpecular(color)
	section.UpdateData(&color[0], offset_specular, UniformBuffer.Std140_size_vec3)
}

func (section *UBOSection) ForceUpdate() {
	direction := section.Direction
	ambient := section.Ambient
	diffuse := section.Diffuse
	specular := section.Specular

	section.UpdateData(&direction[0], offset_direction, UniformBuffer.Std140_size_vec3)
	section.UpdateData(&ambient[0], offset_ambient, UniformBuffer.Std140_size_vec3)
	section.UpdateData(&diffuse[0], offset_diffuse, UniformBuffer.Std140_size_vec3)
	section.UpdateData(&specular[0], offset_specular, UniformBuffer.Std140_size_vec3)

	section.UBOSection.ForceUpdate()
}
