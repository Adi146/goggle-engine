package FrameBuffer

import "github.com/Adi146/goggle-engine/Core/Texture"

type IFrameBuffer interface {
	GetFBO() uint32
	GetType() Type
	GetTextures() []Texture.ITexture
	Destroy()
	Bind()
	Clear()
}
