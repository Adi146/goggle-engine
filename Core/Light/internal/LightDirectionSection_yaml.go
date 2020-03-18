package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"gopkg.in/yaml.v3"
)

func (section *LightDirectionSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightDirection); err != nil {
		return err
	}

	if section.Direction == (GeometryMath.Vector3{}) {
		section.Direction = GeometryMath.Vector3{-1, -1, -1}
	}

	section.ForceUpdate()

	return nil
}
