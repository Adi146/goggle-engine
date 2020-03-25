package PhongShader_test

var (
	vertexShaders = []string{
		"phong.vert",
	}
	fragmentShaders = []string{
		"phong.frag",
		"lights.frag",
		"../../Model/Material/material.frag",
		"../../Light/ShadowMapping/shadow.frag",
	}
	geometryShaders []string
)

//func TestPhongCompile(t *testing.T) {
//	window, _ := TestUtils.CreateTestWindow(t)
//	defer window.Destroy()
//
//	shader, err := PhongShader.NewShaderProgram(vertexShaders, fragmentShaders, geometryShaders)
//	if shader == nil || err != nil {
//		t.Errorf("Error while compiling phong shader %s", err.Error())
//	}
//	defer shader.Destroy()
//}
