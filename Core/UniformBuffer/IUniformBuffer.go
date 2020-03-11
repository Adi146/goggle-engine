package UniformBuffer

import "github.com/Adi146/goggle-engine/Core/Shader"

type IUniformBuffer interface {
	Shader.IUniformBlock

	GetUBO() uint32
	GetType() Type
	Bind()
	UpdateData(data interface{}, offset int, size int)
}
