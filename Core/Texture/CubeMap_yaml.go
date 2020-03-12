package Texture

import "gopkg.in/yaml.v3"

type yamlConfig struct {
	Top    string `yaml:"top"`
	Bottom string `yaml:"bottom"`
	Left   string `yaml:"left"`
	Right  string `yaml:"right"`
	Front  string `yaml:"front"`
	Back   string `yaml:"back"`
}

func (helper *yamlConfig) getArray() []string {
	return []string{
		helper.Right,
		helper.Left,
		helper.Top,
		helper.Bottom,
		helper.Front,
		helper.Back,
	}
}

func (cubeMap *CubeMap) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig yamlConfig
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpCubeMap, err := ImportCubeMap(yamlConfig.getArray(), "")
	if err != nil {
		return err
	}

	*cubeMap = *tmpCubeMap
	return nil
}
