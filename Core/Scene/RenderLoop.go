package Scene

import (
	"github.com/Adi146/goggle-engine/Utils/Log"
)

func RunRenderLoop(scene IScene) {
	if err := scene.Start(); err != nil {
		panic(err)
	}

	for !scene.GetWindow().ShouldClose() {
		scene.GetWindow().PollEvents()

		timeDelta, _ := scene.GetWindow().GetTimeDeltaAndFPS()
		scene.Tick(timeDelta)

		Log.Error(scene.Draw(nil, nil, nil, nil), "Render Error")

		scene.GetWindow().SwapWindow()
	}
}
