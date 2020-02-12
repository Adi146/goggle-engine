package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

type UniformBuffer struct {
	DirectionalLight      `yaml:",inline"`
	ubo.UniformBufferBase `yaml:",inline"`
}

func (buff *UniformBuffer) Init() error {
	buff.Size = 4 * 16

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
	buff.UpdateData(&direction[0], 0, 16)
}

func (buff *UniformBuffer) SetAmbient(color Vector.Vector3) {
	buff.DirectionalLight.SetAmbient(color)
	buff.UpdateData(&color[0], 16, 16)
}

func (buff *UniformBuffer) SetDiffuse(color Vector.Vector3) {
	buff.DirectionalLight.SetDiffuse(color)
	buff.UpdateData(&color[0], 32, 16)
}

func (buff *UniformBuffer) SetSpecular(color Vector.Vector3) {
	buff.DirectionalLight.SetSpecular(color)
	buff.UpdateData(&color[0], 48, 16)
}

func (buff *UniformBuffer) ForceUpdate() {
	buff.UpdateData(&buff.Direction[0], 0, 16)
	buff.UpdateData(&buff.Ambient[0], 16, 16)
	buff.UpdateData(&buff.Diffuse[0], 32, 16)
	buff.UpdateData(&buff.Specular[0], 48, 16)
}
