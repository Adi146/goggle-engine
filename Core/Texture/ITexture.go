package Texture

type ITexture interface {
	GetID() uint32
	GetUnit() *Unit
	GetType() Type
	GetTarget() uint32

	Bind() error
	Unbind()

	SetWrapMode(mode WrapMode)
	GenerateMipMap(lodBias float32)
}
