package PointLight

import (
	"gopkg.in/yaml.v3"
)

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		PointLight PointLight `yaml:",inline"`
	}{
		PointLight: section.PointLight,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	section.SetAmbient(yamlConfig.PointLight.Ambient)
	section.SetDiffuse(yamlConfig.PointLight.Diffuse)
	section.SetSpecular(yamlConfig.PointLight.Specular)
	section.SetLinear(yamlConfig.PointLight.Linear)
	section.SetQuadratic(yamlConfig.PointLight.Quadratic)

	return nil
}
