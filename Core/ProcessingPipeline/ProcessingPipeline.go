package ProcessingPipeline

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type ProcessingPipeline struct {
	Steps  []ProcessingPipelineStep
	Scenes []Scene.IScene
	Window Window.IWindow
}

type ProcessingPipelineStep struct {
	Shader      Shader.IShaderProgram
	FrameBuffer FrameBuffer.IFrameBuffer
	Scene       Scene.IScene
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

		var results []*Model.Texture
		for i := range pipeline.Steps {
			step := pipeline.Steps[len(pipeline.Steps)-1-i]
			step.FrameBuffer.Bind()
			step.Shader.Bind()
			step.Scene.SetActiveShaderProgram(step.Shader)

			for _, result := range results {
				step.Shader.BindObject(result)
			}

			step.FrameBuffer.Clear()
			step.Scene.Draw()

			results = step.FrameBuffer.GetTextures()
		}
		pipeline.Window.SwapWindow()
	}
}
