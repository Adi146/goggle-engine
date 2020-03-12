package Texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Type string

type Texture struct {
	ID     uint32
	Target uint32
	Type   Type
	Unit   *Unit
}

func (texture *Texture) GetUnit() *Unit {
	return texture.Unit
}

func (texture *Texture) GetType() Type {
	return texture.Type
}

func (texture *Texture) Bind() error {
	if texture.Unit == nil {
		unit, err := unitManager.FindUnit(texture)
		if err != nil {
			return err
		}

		texture.Unit = unit
		texture.Unit.Texture = texture
	}

	texture.Unit.Activate()
	gl.BindTexture(texture.Target, texture.ID)

	return nil
}

func (texture *Texture) Unbind() {
	if texture.Unit != nil {
		texture.Unit.Activate()
		gl.BindTexture(texture.Target, 0)

		texture.Unit.Texture = nil
		texture.Unit = nil
	}
}
