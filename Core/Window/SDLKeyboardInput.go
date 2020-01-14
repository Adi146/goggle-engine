package Window

import "github.com/veandco/go-sdl2/sdl"

type SDLKeyboardInput struct {
	keyMap map[string]bool
}

func NewSDLKeyboardInput() *SDLKeyboardInput {
	return &SDLKeyboardInput{
		keyMap: make(map[string]bool),
	}
}

func (input *SDLKeyboardInput) IsKeyPressed(key string) bool {
	val, ok := input.keyMap[key]
	return val && ok
}

func (input *SDLKeyboardInput) pushKeyboardEvent(event *sdl.KeyboardEvent) {
	if event.Type == sdl.KEYDOWN {
		input.keyMap[sdl.GetKeyName(event.Keysym.Sym)] = true
	} else if event.Type == sdl.KEYUP {
		input.keyMap[sdl.GetKeyName(event.Keysym.Sym)] = false
	}
}
