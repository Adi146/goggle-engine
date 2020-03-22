package ProcessingPipeline

import (
	"github.com/Adi146/goggle-engine/Core/Scene"
)

type ProcessingPipeline struct {
	Scene Scene.IScene
}

func (pipeline ProcessingPipeline) Run() {
	for !pipeline.Scene.GetWindow().ShouldClose() {
		pipeline.Scene.GetWindow().PollEvents()

		timeDelta, _ := pipeline.Scene.GetWindow().GetTimeDeltaAndFPS()
		pipeline.Scene.Tick(timeDelta)

		pipeline.Scene.Draw(nil, nil, nil)

		pipeline.Scene.GetWindow().SwapWindow()
	}
}
