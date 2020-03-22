package main

import (
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory"
	"os"
	"runtime"

	_ "github.com/Adi146/goggle-engine/Examples/SceneGraph"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node/CameraNode"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node/LightNode"
	_ "github.com/Adi146/goggle-engine/UI/Control"

	_ "github.com/Adi146/goggle-engine/Core/Shader/PhongShader"

	_ "github.com/ftrvxmtrx/tga"
)

const (
	configFile = "config.yaml"
)

func main() {
	runtime.LockOSThread()

	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scene, err := Factory.ReadConfig(file)
	if err != nil {
		panic(err)
	}

	Scene.RunRenderLoop(scene)
}
