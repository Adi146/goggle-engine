package Shader

type IShaderProgram interface {
	Bind()
	Unbind()

	Destroy()

	BindObject(i interface{}) error
	BindUniform(i interface{}, address string) error
}
