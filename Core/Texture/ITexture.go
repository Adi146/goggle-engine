package Texture

type ITexture interface {
	GetUnit() *Unit

	Bind() error
	Unbind()
}
