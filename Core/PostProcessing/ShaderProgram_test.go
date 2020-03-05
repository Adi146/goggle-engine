package PostProcessing_test

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/TestUtils"
	"testing"
)

var (
	vertexShaders = []string {
		"postProcessing.vert",
	}
	fragmentShaders = [][]string {
		{
			"none.frag",
		},
		{
			"grayscale.frag",
		},
		{
			"inversion.frag",
		},
		{
			"kernel.frag",
		},
	}
)

func TestPostprocessingCompile(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	for _, fragmentShader := range fragmentShaders {
		shader, err := PostProcessing.NewShaderProgram(vertexShaders, fragmentShader)
		if shader == nil || err != nil {
			t.Errorf("Error while compiling postprocessing shader %s", err.Error())
		}
		defer shader.Destroy()
	}
}