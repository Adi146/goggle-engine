package UniformBuffer

type IUniformBufferSection interface {
	ForceUpdate()
	SetUniformBuffer(ubo IUniformBuffer, offset int)
	GetSize() int
}
