package internal

import (
	"gopkg.in/yaml.v3"
)

func (section *LightPositionSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightPosition); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
