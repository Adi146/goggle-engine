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
    `

	if err := yaml.Unmarshal([]byte(data), &light); err != nil {
		t.Errorf("unmarshal failed: %s", err.Error())
	}

	expectedDirection := GeometryMath.Vector3{-1, -1, -1}
	expectedAmbient := GeometryMath.Vector3{0.32, 0.32, 0.32}
	expectedDiffuse := GeometryMath.Vector3{0.8, 0.8, 0.8}
	expectedSpecular := GeometryMath.Vector3{0.8, 0.8, 0.8}

	if light.Direction != expectedDirection {
		t.Errorf("direction value not matching (expecting %f, got %f", light.Direction, expectedDirection)
	}
	if light.Ambient != expectedAmbient {
		t.Errorf("ambient value not matching (expecting %f, got %f", light.Ambient, expectedAmbient)
	}
	if light.Diffuse != expectedDiffuse {
		t.Errorf("diffuse value not matching (expecting %f, got %f", light.Diffuse, expectedDiffuse)
	}
	if light.Specular != expectedSpecular {
		t.Errorf("specular value not matching (expecting %f, got %f", light.Specular, expectedSpecular)
	}

	expectedBinding := uint32(1)

	if light.LightDirectionSection.UniformBuffer == nil || light.LightColorSection.UniformBuffer == nil || light.CameraSection.UniformBuffer == nil {
		t.Errorf("uniform buffer is not set")
	} else {
		if light.LightDirectionSection.UniformBuffer.GetUBO() == 0 || light.LightColorSection.UniformBuffer.GetUBO() == 0 || light.CameraSection.UniformBuffer.GetUBO() == 0 {
			t.Errorf("uniform buffer is not initialized")
		}
		if light.LightDirectionSection.UniformBuffer.GetBinding() != expectedBinding || light.LightColorSection.UniformBuffer.GetBinding() != expectedBinding || light.CameraSection.UniformBuffer.GetBinding() != expectedBinding {
			t.Errorf("uniform buffer binding is not correct")
		}
	}

	expectedProjectionMatrix := *GeometryMath.Orthogonal(-3000, 3000, -3000, 3000, -3000, 3000)
	expectedViewMatrix := *GeometryMath.Identity()

	if !light.ProjectionMatrix.Equals(&expectedProjectionMatrix, 1e-5) {
		t.Errorf("projection matrix not matching (expecting %f, got %f)", expectedProjectionMatrix, light.ProjectionMatrix)
	}
	if light.ViewMatrix != expectedViewMatrix {
		t.Errorf("view matrix not matching (expecting %f, got %f)", expectedViewMatrix, light.ViewMatrix)
	}
}
