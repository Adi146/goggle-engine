package SpotLight

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.SpotLight); err != nil {
		return err
	}

	if section.Direction == (GeometryMath.Vector3{}) {
		section.Direction = GeometryMath.Vector3{0, 0, 1}
	}

	section.ForceUpdate()

	return nil
}
