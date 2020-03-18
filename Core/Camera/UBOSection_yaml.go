package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

func (section *CameraSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.Camera); err != nil {
		return err
	}

	if section.ViewMatrix == (GeometryMath.Matrix4x4{}) {
		section.ViewMatrix = *GeometryMath.Identity()
	}

	section.ForceUpdate()

	return nil
}
