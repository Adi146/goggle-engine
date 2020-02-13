package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	direction_offset = 0
	ambient_offset = 16
	diffuse_offset = 32
	specular_offset = 48
)

type UniformBuffer struct {
	DirectionalLight      `yaml:",inline"`
	ubo.UniformBufferBase `yaml:",inline"`
}

func (buff *UniformBuffer) Init() error {
	buff.Size = 4 * ubo.Std140_size_vec3

	err := buff.UniformBufferBase.Init()
	if err != nil {
		return err
	}

	buff.ForceUpdate()

	return err
}

func (buff *UniformBuffer) Set(light DirectionalLight) {
	buff.DirectionalLight = light
	buff.ForceUpdate()
}

func (buff *UniformBuffer) SetDirection(direction Vector.Vector3) {
	buff.DirectionalLight.SetDiffuse(direction)
	buff.UpdateData(&direction[0], direction_offset, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetAmbient(color Vector.Vector3) {
	buff.DirectionalLight.SetAmbient(color)
	buff.UpdateData(&color[0], ambient_offset, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetDiffuse(color Vector.Vector3) {
	buff.DirectionalLight.SetDiffuse(color)
	buff.UpdateData(&color[0], diffuse_offset, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) SetSpecular(color Vector.Vector3) {
	buff.DirectionalLight.SetSpecular(color)
	buff.UpdateData(&color[0], specular_offset, ubo.Std140_size_vec3)
}

func (buff *UniformBuffer) ForceUpdate() {
	buff.UpdateData(&buff.Direction[0], direction_offset, ubo.Std140_size_vec3)
	buff.UpdateData(&buff.Ambient[0], ambient_offset, ubo.Std140_size_vec3)
	buff.UpdateData(&buff.Diffuse[0], diffuse_offset, ubo.Std140_size_vec3)
	buff.UpdateData(&buff.Specular[0], specular_offset, ubo.Std140_size_vec3)
}
