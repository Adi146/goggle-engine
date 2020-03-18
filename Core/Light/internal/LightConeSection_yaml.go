package internal

import "gopkg.in/yaml.v3"

func (section *LightConeSection) UnmarshalYAML(value *yaml.Node) error {
	if err := value.Decode(&section.LightCone); err != nil {
		return err
	}
	section.ForceUpdate()

	return nil
}
