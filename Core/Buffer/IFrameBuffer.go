package Buffer

type IFrameBuffer interface {
	Destroy()

	GetSize() (int32, int32)
	GetFBO() uint32
}
