package Window

type IMouseInput interface {
	GetRelativeMovement() (float32, float32)
}
