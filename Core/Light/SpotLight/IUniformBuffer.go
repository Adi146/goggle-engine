package SpotLight

import ubo "github.com/Adi146/goggle-engine/Core/UniformBuffer"

type IUniformBuffer interface {
	ubo.IUniformBuffer
	GetNewElement() (*UniformBufferElement, error)
}
