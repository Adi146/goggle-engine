package Light_test

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Utils/TestUtils"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestUBODirectionalLight_UnmarshalYAML(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()

	var light Light.UBODirectionalLight
	var data = `
directionalLight:
    direction: [-1, -1, -1]
    ambient: [0.32, 0.32, 0.32]
    diffuse: [0.8, 0.8, 0.8]
    specular: [0.8, 0.8, 0.8]
uniformBuffer:
    binding: 1
shadowMap:
    distance: 200
    transitionDistance: 20
    frameBuffer:
        width: 4096
        height: 4096
    `

	if err := yaml.Unmarshal([]byte(data), &light); err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	}

	expectedDirection := GeometryMath.Vector3{-1, -1, -1}
	expectedAmbient := GeometryMath.Vector3{0.32, 0.32, 0.32}
	expectedDiffuse := GeometryMath.Vector3{0.8, 0.8, 0.8}
	expectedSpecular := GeometryMath.Vector3{0.8, 0.8, 0.8}

	if light.Direction.Get() != expectedDirection {
		t.Errorf("direction value not matching (expecting %f, got %f", expectedDirection, light.Direction.Get())
	}
	if light.Ambient.Get() != expectedAmbient {
		t.Errorf("ambient value not matching (expecting %f, got %f", expectedAmbient, light.Ambient.Get())
	}
	if light.Diffuse.Get() != expectedDiffuse {
		t.Errorf("diffuse value not matching (expecting %f, got %f", expectedDiffuse, light.Diffuse.Get())
	}
	if light.Specular.Get() != expectedSpecular {
		t.Errorf("specular value not matching (expecting %f, got %f", expectedSpecular, light.Specular.Get())
	}

	expectedDistance := float32(200)
	expectedTransitionDistance := float32(20)

	if light.ShadowMap.Distance.Get() != expectedDistance {
		t.Errorf("distance value not matching (expecting %f, got %f)",expectedDistance , light.ShadowMap.Distance.Get())
	}
	if light.ShadowMap.TransitionDistance.Get() != expectedTransitionDistance {
		t.Errorf("transition distance value not matching (expecting %f, got %f)",expectedTransitionDistance , light.ShadowMap.TransitionDistance.Get())
	}
}
