package Texture

type ITexture interface {
	GetUnit() *Unit
	GetType() Type

	Bind() error
	Unbind()
}
