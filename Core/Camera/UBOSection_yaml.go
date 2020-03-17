package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Camera Camera `yaml:",inline"`
	}{
		Camera: section.Camera,
	}

	if yamlConfig.Camera.ViewMatrix == (GeometryMath.Matrix4x4{}) {
		yamlConfig.Camera.ViewMatrix = *GeometryMath.Identity()
	}

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	section.SetProjectionMatrix(yamlConfig.Camera.ProjectionMatrix)
	section.SetViewMatrix(yamlConfig.Camera.ViewMatrix)

	return nil
}
