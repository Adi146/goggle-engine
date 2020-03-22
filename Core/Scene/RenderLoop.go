package Scene

func RunRenderLoop(scene IScene) {
	for !scene.GetWindow().ShouldClose() {
		scene.GetWindow().PollEvents()

		timeDelta, _ := scene.GetWindow().GetTimeDeltaAndFPS()
		scene.Tick(timeDelta)

		scene.Draw(nil, nil, nil)

		scene.GetWindow().SwapWindow()
	}
}
