package Texture

type ITexture interface {
	Bind(unit uint32)
	Unbind()
}
