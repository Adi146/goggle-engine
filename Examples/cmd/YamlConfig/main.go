package main

import (
	"os"
	"runtime"

	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"

	_ "github.com/Adi146/goggle-engine/Examples/SceneGraph"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node/LightNode"
	_ "github.com/Adi146/goggle-engine/UI/Control"

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

	config, err := YamlFactory.ReadConfig(file)
	if err != nil {
		panic(err)
	}

	config.Pipeline.Run()
}
