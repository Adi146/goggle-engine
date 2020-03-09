package Scene

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type ProcessingPipeline struct {
	Steps  []ProcessingPipelineStep
	Scenes []IScene
	Window Window.IWindow
}

type ProcessingPipelineStep struct {
	FrameBuffer    FrameBuffer.IFrameBuffer
	Scene          IScene
	EnforcedShader Shader.IShaderProgram
}

func (pipeline ProcessingPipeline) Run() {
	for _, scene := range pipeline.Scenes {
		scene.SetKeyboardInput(pipeline.Window.GetKeyboardInput())
		scene.SetMouseInput(pipeline.Window.GetMouseInput())
	}

	for !pipeline.Window.ShouldClose() {
		pipeline.Window.PollEvents()

		timeDelta, _ := pipeline.Window.GetTimeDeltaAndFPS()
		for _, scene := range pipeline.Scenes {
			scene.Tick(timeDelta)
		}

		for i := range pipeline.Steps {
			step := pipeline.Steps[len(pipeline.Steps)-1-i]
			step.FrameBuffer.Bind()

			step.FrameBuffer.Clear()
			step.Scene.Draw(&step)
		}

		pipeline.Window.SwapWindow()
	}
}
