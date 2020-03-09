package UniformBuffer

type IUniformBuffer interface {
	Init() error
	GetUBO() uint32
	Bind()
	UpdateData(data interface{}, offset int, size int)
	GetBinding() uint32

	ForceUpdate()
}
