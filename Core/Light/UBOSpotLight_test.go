package Light_test

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestUBOSpotLight_UnmarshalYAML(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	lightStruct := struct {
		Lights []Light.UBOSpotLight `yaml:"lights"`
	}{}

	var data = `
lights:
    - spotLight:
          innerCone: 0.80
          outerCone: 0.95
          ambient: [0.1, 0.1, 0]
          diffuse: [1, 1, 0]
          specular: [1, 1, 0]
          linear: 0.0014
          quadratic: 0.000007
          direction: [1, 0, 1]
      uniformBuffer: spotLight
    - spotLight:
          innerCone: 0.80
          outerCone: 0.95
          ambient: [0, 0.1, 0]
          diffuse: [0, 1, 0]
          specular: [0, 1, 0]
          linear: 0.0014
          quadratic: 0.000007
          direction: [0, 0, 1]
      uniformBuffer:
          binding: 3
          id: spotLight
    `

	if err := yaml.Unmarshal([]byte(data), &lightStruct); err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	}

	type expectedResult struct {
		Direction GeometryMath.Vector3
		Position  GeometryMath.Vector3
		Ambient   GeometryMath.Vector3
		Diffuse   GeometryMath.Vector3
		Specular  GeometryMath.Vector3
		Linear    float32
		Quadratic float32
		InnerCone float32
		OuterCone float32
		Binding   uint32
	}

	expectedResults := []expectedResult{
		{
			Direction: GeometryMath.Vector3{1, 0, 1},
			Position:  GeometryMath.Vector3{0, 0, 0},
			Ambient:   GeometryMath.Vector3{0.1, 0.1, 0.0},
			Diffuse:   GeometryMath.Vector3{1.0, 1.0, 0.0},
			Specular:  GeometryMath.Vector3{1.0, 1.0, 0.0},
			Linear:    0.0014,
			Quadratic: 0.000007,
			InnerCone: 0.80,
			OuterCone: 0.95,
			Binding:   3,
		},
		{
			Direction: GeometryMath.Vector3{0, 0, 1},
			Position:  GeometryMath.Vector3{0, 0, 0},
			Ambient:   GeometryMath.Vector3{0.0, 0.1, 0.0},
			Diffuse:   GeometryMath.Vector3{0.0, 1.0, 0.0},
			Specular:  GeometryMath.Vector3{0.0, 1.0, 0.0},
			Linear:    0.0014,
			Quadratic: 0.000007,
			InnerCone: 0.80,
			OuterCone: 0.95,
			Binding:   3,
		},
	}

	for i, light := range lightStruct.Lights {
		if light.Direction.Get() != expectedResults[i].Direction {
			t.Errorf("[light %d] direction value not matching (expecting %f, got %f)", i, light.Direction.Get(), expectedResults[i].Direction)
		}
		if light.Position.Get() != expectedResults[i].Position {
			t.Errorf("[light %d] position value not matching (expecting %f, got %f)", i, light.Position.Get(), expectedResults[i].Position)
		}
		if light.Ambient.Get() != expectedResults[i].Ambient {
			t.Errorf("[light %d] ambient value not matching (expecting %f, got %f)", i, light.Ambient.Get(), expectedResults[i].Ambient)
		}
		if light.Diffuse.Get() != expectedResults[i].Diffuse {
			t.Errorf("[light %d] diffuse value not matching (expecting %f, got %f)", i, light.Diffuse.Get(), expectedResults[i].Diffuse)
		}
		if light.Specular.Get() != expectedResults[i].Specular {
			t.Errorf("[light %d] specular value not matching (expecting %f, got %f)", i, light.Specular.Get(), expectedResults[i].Specular)
		}
		if light.Linear.Get() != expectedResults[i].Linear {
			t.Errorf("[light %d] linear value not matching (expecting %f, got %f)", i, light.Linear.Get(), expectedResults[i].Linear)
		}
		if light.Quadratic.Get() != expectedResults[i].Quadratic {
			t.Errorf("[light %d] quadratic value not matching (expecting %f, got %f)", i, light.Quadratic.Get(), expectedResults[i].Quadratic)
		}
		if light.InnerCone.Get() != expectedResults[i].InnerCone {
			t.Errorf("[light %d] innerCone value not matching (expecting %f, got %f)", i, light.InnerCone.Get(), expectedResults[i].InnerCone)
		}
		if light.OuterCone.Get() != expectedResults[i].OuterCone {
			t.Errorf("[light %d] outerCone value not matching (expecting %f, got %f)", i, light.OuterCone.Get(), expectedResults[i].OuterCone)
		}
	}
}
