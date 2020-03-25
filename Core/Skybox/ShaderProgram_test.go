package Skybox_test

import (
	"github.com/Adi146/goggle-engine/Core/Skybox"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"testing"
)

var (
	vertexShaders = []string{
		"skybox.vert",
	}
	fragmentShaders = []string{
		"skybox.frag",
	}
	geometryShaders []string
)

func TestSkyboxCompile(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	shader, err := Skybox.NewShaderProgram(vertexShaders, fragmentShaders, geometryShaders)
	if shader == nil || err != nil {
		t.Errorf("Error while compiling skybox shader %s", err.Error())
	}
	defer shader.Destroy()
}
