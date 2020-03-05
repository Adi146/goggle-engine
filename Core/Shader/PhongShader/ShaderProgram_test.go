package PhongShader_test

import (
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/Core/TestUtils"
	"testing"
)

var (
	vertexShaders = []string {
		"phong.vert",
	}
	fragmentShaders = []string {
		"phong.frag",
		"material.frag",
		"lights.frag",
	}
)

func TestCompile(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	shader, err := PhongShader.NewPhongShaderProgram(vertexShaders, fragmentShaders)
	if shader == nil || err != nil {
		t.Errorf("Error while compiling phong shader %s", err.Error())
	}
	defer shader.Destroy()
}