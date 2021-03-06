package Light_test

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestUBOPointLight_UnmarshalYAML(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	lightStruct := struct {
		Lights []Light.PointLight `yaml:"lights"`
	}{}

	var data = `
lights:
    - pointLight:
          ambient: [0.2, 0.0, 0.0]
          diffuse: [1.0, 0.0, 0.0]
          specular: [1.0, 0.0, 0.0]
          linear: 0.0014
          quadratic: 0.000007
      uniformBuffer:
          id: pointLight
          binding: 2
      shadowMap:
          projection:
              perspective:
                  fovy: 90
                  aspect: 1
                  near: 1
                  far: 3250
          frameBuffer:
              width: 1024
              height: 1024
    - pointLight:
          ambient: [0.0, 0.0, 0.2]
          diffuse: [0.0, 0.0, 1.0]
          specular: [0.0, 0.0, 1.0]
          linear: 0.0014
          quadratic: 0.000007
      uniformBuffer: pointLight
      shadowMap:
          distance: 3250
          frameBuffer:
              width: 1024
              height: 1024
    `

	if err := yaml.Unmarshal([]byte(data), &lightStruct); err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	}

	type expectedResult struct {
		Position  GeometryMath.Vector3
		Ambient   GeometryMath.Vector3
		Diffuse   GeometryMath.Vector3
		Specular  GeometryMath.Vector3
		Linear    float32
		Quadratic float32
		Binding   uint32
	}

	expectedResults := []expectedResult{
		{
			Position:  GeometryMath.Vector3{0, 0, 0},
			Ambient:   GeometryMath.Vector3{0.2, 0.0, 0.0},
			Diffuse:   GeometryMath.Vector3{1.0, 0.0, 0.0},
			Specular:  GeometryMath.Vector3{1.0, 0.0, 0.0},
			Linear:    0.0014,
			Quadratic: 0.000007,
			Binding:   2,
		},
		{
			Position:  GeometryMath.Vector3{0, 0, 0},
			Ambient:   GeometryMath.Vector3{0.0, 0.0, 0.2},
			Diffuse:   GeometryMath.Vector3{0.0, 0.0, 1.0},
			Specular:  GeometryMath.Vector3{0.0, 0.0, 1.0},
			Linear:    0.0014,
			Quadratic: 0.000007,
			Binding:   2,
		},
	}

	for i, light := range lightStruct.Lights {
		if light.PointLight.Position != expectedResults[i].Position {
			t.Errorf("[light %d] position value not matching (expecting %f, got %f)", i, expectedResults[i].Position, light.PointLight.Position)
		}
		if light.PointLight.Ambient != expectedResults[i].Ambient {
			t.Errorf("[light %d] ambient value not matching (expecting %f, got %f)", i, expectedResults[i].Ambient, light.PointLight.Ambient)
		}
		if light.PointLight.Diffuse != expectedResults[i].Diffuse {
			t.Errorf("[light %d] diffuse value not matching (expecting %f, got %f)", i, expectedResults[i].Diffuse, light.PointLight.Diffuse)
		}
		if light.PointLight.Specular != expectedResults[i].Specular {
			t.Errorf("[light %d] specular value not matching (expecting %f, got %f)", i, expectedResults[i].Specular, light.PointLight.Specular)
		}
		if light.PointLight.Linear != expectedResults[i].Linear {
			t.Errorf("[light %d] linear value not matching (expecting %f, got %f)", i, expectedResults[i].Linear, light.PointLight.Linear)
		}
		if light.PointLight.Quadratic != expectedResults[i].Quadratic {
			t.Errorf("[light %d] quadratic value not matching (expecting %f, got %f)", i, expectedResults[i].Quadratic, light.PointLight.Quadratic)
		}
	}
}
