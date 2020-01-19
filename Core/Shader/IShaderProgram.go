package Shader

type IShaderProgram interface {
	Bind()
	Unbind()

	Destroy()

	BeginDraw() []error
	EndDraw()

	BindObject(i interface{}) []error
}
