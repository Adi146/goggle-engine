package internal

import (
	"gopkg.in/yaml.v3"
)

func (section *LightColorSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightColor); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
