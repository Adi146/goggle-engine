package Window

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	typeFlags = map[string]uint32{
		"window":     0,
		"borderless": sdl.WINDOW_BORDERLESS,
		"fullscreen": sdl.WINDOW_FULLSCREEN,
	}

	syncFlags = map[string]int{
		"normal":   0,
		"adaptive": -1,
		"vertical": 1,
	}
)

type SDLWindow struct {
	FrameBuffer.FrameBufferBase `yaml:",inline"`

	window    *sdl.Window
	glContext sdl.GLContext

	keyboardInput *SDLKeyboardInput
	mouseInput    *SDLMouseInput

	performanceCounterFrequency uint64
	lastCounter                 uint64
	shouldClose                 bool

	Title string `yaml:"title"`
	Type  string `yaml:"type"`
	Sync  string `yaml:"sync"`
}

func (window *SDLWindow) Init() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	sdl.GLSetAttribute(sdl.GL_RED_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_BUFFER_SIZE, 32)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	sdl.SetRelativeMouseMode(true)

	sdlWindow, err := sdl.CreateWindow(
		window.Title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		window.Width,
		window.Height,
		sdl.WINDOW_OPENGL|typeFlags[window.Type],
	)
	if err != nil {
		return err
	}

	glContext, err := sdlWindow.GLCreateContext()
	if err != nil {
		return err
	}
	if err := gl.Init(); err != nil {
		return err
	}

	sdl.GLSetSwapInterval(syncFlags[window.Sync])

	window.window = sdlWindow
	window.glContext = glContext
	window.keyboardInput = NewSDLKeyboardInput()
	window.mouseInput = NewSDLMouseInput()
	window.performanceCounterFrequency = sdl.GetPerformanceFrequency()
	window.lastCounter = sdl.GetPerformanceCounter()

	return nil
}

func (window *SDLWindow) Destroy() {
	sdl.GLDeleteContext(window.glContext)
	window.window.Destroy()
	sdl.Quit()
}

func (window *SDLWindow) PollEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			window.shouldClose = true
			break
		case *sdl.KeyboardEvent:
			window.keyboardInput.pushKeyboardEvent(t)
		case *sdl.MouseMotionEvent:
			window.mouseInput.pushMouseMotionEvent(t)
		case *sdl.MouseButtonEvent:
			window.mouseInput.pushMouseButtonEvent(t)
		case *sdl.MouseWheelEvent:
			window.mouseInput.pushMouseWheelEvent(t)
		}
	}
}

func (window *SDLWindow) SwapWindow() {
	window.window.GLSwap()
}

func (window *SDLWindow) GetTimeDeltaAndFPS() (float32, uint32) {
	endCounter := sdl.GetPerformanceCounter()
	counterElapsed := endCounter - window.lastCounter
	timeDelta := (float32)(counterElapsed) / (float32)(window.performanceCounterFrequency)
	fps := (uint32)((float32)(window.performanceCounterFrequency) / (float32)(counterElapsed))
	window.lastCounter = endCounter

	return timeDelta, fps
}

func (window *SDLWindow) ShouldClose() bool {
	return window.shouldClose
}

func (window *SDLWindow) GetSize() (int32, int32) {
	return window.window.GetSize()
}

func (window *SDLWindow) GetKeyboardInput() IKeyboardInput {
	return window.keyboardInput
}

func (window *SDLWindow) GetMouseInput() IMouseInput {
	return window.mouseInput
}

func (window *SDLWindow) GetTextures() []*Texture.Texture {
	return nil
}
