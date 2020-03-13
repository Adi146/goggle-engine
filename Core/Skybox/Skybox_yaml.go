package Skybox

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"
)

func (skybox *Skybox) UnmarshalYAML(value *yaml.Node) error {
	var texture Texture.CubeMap
	if err := value.Decode(&texture); err != nil {
		return err
	}
	texture.Type = Texture.SkyboxTexture

	tmpSkybox, err := NewSkybox(&texture)
	if err != nil {
		return err
	}

	*skybox = *tmpSkybox
	return nil
}
