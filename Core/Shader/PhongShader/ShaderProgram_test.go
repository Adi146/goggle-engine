package PhongShader_test

import (
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"testing"
)

var (
	vertexShaders = []string{
		"phong.vert",
	}
	fragmentShaders = []string{
		"phong.frag",
		"lights.frag",
		"../../Model/Material/material.frag",
		"../../Shadow/shadow.frag",
	}
)

func TestPhongCompile(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	shader, err := PhongShader.NewShaderProgram(vertexShaders, fragmentShaders)
	if shader == nil || err != nil {
		t.Errorf("Error while compiling phong shader %s", err.Error())
	}
	defer shader.Destroy()
}
