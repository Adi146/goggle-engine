package DirectionalLight

import (
	"gopkg.in/yaml.v3"
)

func (section *UBOSection) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		DirectionalLight DirectionalLight `yaml:",inline"`
	}{
		DirectionalLight: section.DirectionalLight,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	section.SetDirection(yamlConfig.DirectionalLight.Direction)
	section.SetAmbient(yamlConfig.DirectionalLight.Ambient)
	section.SetDiffuse(yamlConfig.DirectionalLight.Diffuse)
	section.SetSpecular(yamlConfig.DirectionalLight.Specular)

	return nil
}
