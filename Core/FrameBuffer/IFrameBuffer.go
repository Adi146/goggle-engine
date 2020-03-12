package FrameBuffer

import "github.com/Adi146/goggle-engine/Core/Texture"

type IFrameBuffer interface {
	Init() error
	Destroy()

	Bind()
	Unbind()

	GetSize() (int32, int32)
	GetFBO() uint32

	Clear()

	GetTextures() []*Texture.Texture
}
