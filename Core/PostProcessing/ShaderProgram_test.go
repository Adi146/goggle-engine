package PostProcessing_test

import (
	"github.com/Adi146/goggle-engine/Core/PostProcessing"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"testing"
)

var (
	vertexShaders = []string{
		"postProcessing.vert",
	}
	fragmentShaders = [][]string{
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
	geometryShaders []string
)

func TestPostprocessingCompile(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	for _, fragmentShader := range fragmentShaders {
		shader, err := PostProcessing.NewShaderProgram(vertexShaders, fragmentShader, geometryShaders)
		if shader == nil || err != nil {
			t.Errorf("Error while compiling postprocessing shader %s [vertexShaders: %s, fragmentShaders: %s]", err.Error(), vertexShaders, fragmentShader)
		}
		shader.Destroy()
	}
}
