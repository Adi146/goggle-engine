package ProcessingPipeline

import (
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type ProcessingPipeline struct {
	Steps  []Step
	Scenes []Scene.IScene
	Window Window.IWindow
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
			step.Execute()
		}

		pipeline.Window.SwapWindow()
	}
}
