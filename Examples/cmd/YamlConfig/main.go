package main

import (
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/YamlFactory"
	"os"
	"runtime"

	_ "github.com/Adi146/goggle-engine/Examples/SceneGraph"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node"
	_ "github.com/Adi146/goggle-engine/SceneGraph/Node/LightNode"
	_ "github.com/Adi146/goggle-engine/UI/Control"
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
