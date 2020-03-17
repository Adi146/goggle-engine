package PointLight

import (
	"gopkg.in/yaml.v3"
)

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.PointLight); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
