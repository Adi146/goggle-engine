package Window

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/veandco/go-sdl2/sdl"
	"gopkg.in/yaml.v3"

	"github.com/go-gl/gl/v4.3-core/gl"
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

	avgFps       float32
	FpsSmoothing float32

	title       string
	TitlebarFPS bool
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
		FpsSmoothing:                0.99,
		title:                       title,
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
	window.avgFps = (window.avgFps * window.FpsSmoothing) + (float32(fps) * (1.0 - window.FpsSmoothing))

	if window.TitlebarFPS {
		window.Window.SetTitle(window.title + fmt.Sprintf(" [Average FPS: %.2f]", window.avgFps))
	}

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

func (window *SDLWindow) GetTitle() string {
	return window.title
}

func (window *SDLWindow) SetTitle(title string) {
	window.title = title
	window.Window.SetTitle(title)
}

func (window *SDLWindow) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Viewport    FrameBuffer.Viewport `yaml:",inline"`
		Title       string               `yaml:"title"`
		TitlebarFPS bool                 `yaml:"titlebarFPS"`
		Sync        SyncInterval         `yaml:"sync"`
		Flags       Flag                 `yaml:"flags"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	sdlWindow, err := NewSDLWindow(yamlConfig.Viewport, yamlConfig.Title, yamlConfig.Flags)
	if err != nil {
		return err
	}

	sdlWindow.TitlebarFPS = yamlConfig.TitlebarFPS

	*window = *sdlWindow
	sdl.GLSetSwapInterval((int)(yamlConfig.Sync))

	return nil
}
