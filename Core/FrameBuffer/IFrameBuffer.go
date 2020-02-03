package FrameBuffer

import (
	"github.com/Adi146/goggle-engine/Core/Model"
)

type IFrameBuffer interface {
	Init() error
	Destroy()

	Bind()

	GetSize() (int32, int32)
	GetFBO() uint32

	Clear()

	GetTextures() []*Model.Texture
}
