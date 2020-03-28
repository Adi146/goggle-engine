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
    projection:
        orthogonal:
            left: -3000
            right: 3000
            bottom: -3000
            top: 3000
            near: -3000
            far: 3000
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
		t.Errorf("direction value not matching (expecting %f, got %f", light.Direction.Get(), expectedDirection)
	}
	if light.Ambient.Get() != expectedAmbient {
		t.Errorf("ambient value not matching (expecting %f, got %f", light.Ambient.Get(), expectedAmbient)
	}
	if light.Diffuse.Get() != expectedDiffuse {
		t.Errorf("diffuse value not matching (expecting %f, got %f", light.Diffuse.Get(), expectedDiffuse)
	}
	if light.Specular.Get() != expectedSpecular {
		t.Errorf("specular value not matching (expecting %f, got %f", light.Specular.Get(), expectedSpecular)
	}

	expectedProjectionMatrix := *GeometryMath.Orthogonal(-3000, 3000, -3000, 3000, -3000, 3000)

	if !light.ShadowMap.Projection.Equals(&expectedProjectionMatrix, 1e-5) {
		t.Errorf("projection matrix not matching (expecting %f, got %f)", expectedProjectionMatrix, light.ShadowMap.Projection)
	}
}
