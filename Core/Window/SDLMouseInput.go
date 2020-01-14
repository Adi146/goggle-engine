package Window

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type SDLMouseInput struct {
	xRel int32
	yRel int32
}

func NewSDLMouseInput() *SDLMouseInput {
	return &SDLMouseInput{}
}

func (input *SDLMouseInput) GetRelativeMovement() (float32, float32) {
	xRel := input.xRel
	yRel := input.yRel

	input.xRel = 0
	input.yRel = 0

	return float32(xRel), float32(yRel)
}

func (input *SDLMouseInput) pushMouseMotionEvent(event *sdl.MouseMotionEvent) {
	input.xRel += event.XRel
	input.yRel += event.YRel
}

func (input *SDLMouseInput) pushMouseButtonEvent(event *sdl.MouseButtonEvent) {
	fmt.Printf("mouse button (X: %d, Y: %d) \n", event.X, event.Y)
}

func (input *SDLMouseInput) pushMouseWheelEvent(event *sdl.MouseWheelEvent) {
	fmt.Printf("mouse wheel (X: %d, Y: %d) \n", event.X, event.Y)
}
