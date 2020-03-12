package Texture

import (
	"gopkg.in/yaml.v3"
)

func (texture *Texture2D) UnmarshalYAML(value *yaml.Node) error {
	var filename string
	if err := value.Decode(&filename); err != nil {
		return err
	}

	tmpTexture, err := ImportTexture(filename, "")
	if err != nil {
		return err
	}

	*texture = *tmpTexture
	return nil
}
