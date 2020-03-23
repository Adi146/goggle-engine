package FrameBuffer

type IFrameBuffer interface {
	GetFBO() uint32
	GetType() Type
	Destroy()
	Bind()
	Clear()
}
