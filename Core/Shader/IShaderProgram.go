package Shader

type IShaderProgram interface {
	Bind()
	Unbind()

	Destroy()

	GetUniformAddress(i interface{}) (string, error)

	BindObject(i interface{}) error
	BindUniform(i interface{}, address string) error
}
