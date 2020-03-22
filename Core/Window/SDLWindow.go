package Window

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/veandco/go-sdl2/sdl"
	"gopkg.in/yaml.v3"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type SDLWindow struct {
	FrameBuffer.FrameBuffer
	*sdl.Window

	glContext sdl.GLContext

	keyboardInput *SDLKeyboardInput
	mouseInput    *SDLMouseInput

	performanceCounterFrequency uint64
	lastCounter                 uint64
	shouldClose                 bool
}

func NewSDLWindow(viewport FrameBuffer.Viewport, title string, flags Flag) (*SDLWindow, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	sdl.GLSetAttribute(sdl.GL_RED_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, 8)
	sdl.GLSetAttribute(sdl.GL_BUFFER_SIZE, 32)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	sdl.SetRelativeMouseMode(true)

	sdlWindow, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		viewport.Width,
		viewport.Height,
		sdl.WINDOW_OPENGL|uint32(flags),
	)
	if err != nil {
		return nil, err
	}

	glContext, err := sdlWindow.GLCreateContext()
	if err != nil {
		return nil, err
	}
	if err := gl.Init(); err != nil {
		return nil, err
	}

	return &SDLWindow{
		FrameBuffer: FrameBuffer.FrameBuffer{
			Viewport: viewport,
		},
		Window:                      sdlWindow,
		glContext:                   glContext,
		keyboardInput:               NewSDLKeyboardInput(),
		mouseInput:                  NewSDLMouseInput(),
		performanceCounterFrequency: sdl.GetPerformanceFrequency(),
		lastCounter:                 sdl.GetPerformanceCounter(),
	}, nil
}

func (window *SDLWindow) Destroy() {
	sdl.GLDeleteContext(window.glContext)
	window.Window.Destroy()
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
	window.Window.GLSwap()
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

func (window *SDLWindow) GetKeyboardInput() IKeyboardInput {
	return window.keyboardInput
}

func (window *SDLWindow) GetMouseInput() IMouseInput {
	return window.mouseInput
}

func (window *SDLWindow) GetTextures() []Texture.ITexture {
	return nil
}

func (window *SDLWindow) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Viewport FrameBuffer.Viewport `yaml:",inline"`
		Title    string               `yaml:"title"`
		Sync     SyncInterval         `yaml:"sync"`
		Flags    Flag                 `yaml:"flags"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	sdlWindow, err := NewSDLWindow(yamlConfig.Viewport, yamlConfig.Title, yamlConfig.Flags)
	if err != nil {
		return err
	}

	*window = *sdlWindow
	sdl.GLSetSwapInterval((int)(yamlConfig.Sync))

	return nil
}
