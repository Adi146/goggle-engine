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

	var light Light.DirectionalLight
	var data = `
directionalLight:
    direction: [-1, -1, -1]
    ambient: [0.32, 0.32, 0.32]
    diffuse: [0.8, 0.8, 0.8]
    specular: [0.8, 0.8, 0.8]
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

	if light.DirectionalLight.Direction != expectedDirection {
		t.Errorf("direction value not matching (expecting %f, got %f", expectedDirection, light.DirectionalLight.Direction)
	}
	if light.DirectionalLight.Ambient != expectedAmbient {
		t.Errorf("ambient value not matching (expecting %f, got %f", expectedAmbient, light.DirectionalLight.Ambient)
	}
	if light.DirectionalLight.Diffuse != expectedDiffuse {
		t.Errorf("diffuse value not matching (expecting %f, got %f", expectedDiffuse, light.DirectionalLight.Diffuse)
	}
	if light.DirectionalLight.Specular != expectedSpecular {
		t.Errorf("specular value not matching (expecting %f, got %f", expectedSpecular, light.DirectionalLight.Specular)
	}

	expectedDistance := float32(200)
	expectedTransitionDistance := float32(20)

	if light.ShadowMap.Distance != expectedDistance {
		t.Errorf("distance value not matching (expecting %f, got %f)", expectedDistance, light.ShadowMap.Distance)
	}
	if light.ShadowMap.TransitionDistance != expectedTransitionDistance {
		t.Errorf("transition distance value not matching (expecting %f, got %f)", expectedTransitionDistance, light.ShadowMap.TransitionDistance)
	}
}
