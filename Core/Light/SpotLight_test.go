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
		Lights []Light.SpotLight `yaml:"lights"`
	}{}

	var data = `
lights:
    - spotLight:
          innerCone: 12.5
          outerCone: 17.5
          ambient: [0.1, 0.1, 0]
          diffuse: [1, 1, 0]
          specular: [1, 1, 0]
          linear: 0.0014
          quadratic: 0.000007
          direction: [1, 0, 1]
      uniformBuffer: spotLight
      shadowMap:
          distance: 3250
          frameBuffer:
              width: 1024
              height: 1024
    - spotLight:
          innerCone: 20
          outerCone: 45
          ambient: [0, 0.1, 0]
          diffuse: [0, 1, 0]
          specular: [0, 1, 0]
          linear: 0.0014
          quadratic: 0.000007
          direction: [0, 0, 1]
      uniformBuffer:
          binding: 3
          id: spotLight
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
			InnerCone: 0.976296,
			OuterCone: 0.953717,
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
			InnerCone: 0.939693,
			OuterCone: 0.707107,
			Binding:   3,
		},
	}

	for i, light := range lightStruct.Lights {
		if light.SpotLight.Direction.Get() != expectedResults[i].Direction {
			t.Errorf("[light %d] direction value not matching (expecting %f, got %f)", i, expectedResults[i].Direction, light.SpotLight.Direction.Get())
		}
		if light.SpotLight.Position.Get() != expectedResults[i].Position {
			t.Errorf("[light %d] position value not matching (expecting %f, got %f)", i, expectedResults[i].Position, light.SpotLight.Position.Get())
		}
		if light.SpotLight.Ambient.Get() != expectedResults[i].Ambient {
			t.Errorf("[light %d] ambient value not matching (expecting %f, got %f)", i, expectedResults[i].Ambient, light.SpotLight.Ambient.Get())
		}
		if light.SpotLight.Diffuse.Get() != expectedResults[i].Diffuse {
			t.Errorf("[light %d] diffuse value not matching (expecting %f, got %f)", i, expectedResults[i].Diffuse, light.SpotLight.Diffuse.Get())
		}
		if light.SpotLight.Specular.Get() != expectedResults[i].Specular {
			t.Errorf("[light %d] specular value not matching (expecting %f, got %f)", i, expectedResults[i].Specular, light.SpotLight.Specular.Get())
		}
		if light.SpotLight.Linear.Get() != expectedResults[i].Linear {
			t.Errorf("[light %d] linear value not matching (expecting %f, got %f)", i, expectedResults[i].Linear, light.SpotLight.Linear.Get())
		}
		if light.SpotLight.Quadratic.Get() != expectedResults[i].Quadratic {
			t.Errorf("[light %d] quadratic value not matching (expecting %f, got %f)", i, expectedResults[i].Quadratic, light.SpotLight.Quadratic.Get())
		}
		if !GeometryMath.Equals(light.SpotLight.InnerCone.Get(), expectedResults[i].InnerCone, 1e-5) {
			t.Errorf("[light %d] innerCone value not matching (expecting %f, got %f)", i, expectedResults[i].InnerCone, light.SpotLight.InnerCone.Get())
		}
		if !GeometryMath.Equals(light.SpotLight.OuterCone.Get(), expectedResults[i].OuterCone, 1e-5) {
			t.Errorf("[light %d] outerCone value not matching (expecting %f, got %f)", i, expectedResults[i].OuterCone, light.SpotLight.OuterCone.Get())
		}
	}
}
