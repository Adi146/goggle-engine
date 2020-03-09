package Texture

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var unitManager UnitManager

type UnitManager []*Unit

func (man *UnitManager) populate() {
	if len(*man) == 0 {
		var numTextureUnits int32
		gl.GetIntegerv(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS, &numTextureUnits)

		*man = make([]*Unit, numTextureUnits)
		for i := int32(0); i < numTextureUnits; i++ {
			(*man)[i] = &Unit{
				ID: uint32(i),
				Texture: nil,
			}
		}
	}
}

func (man *UnitManager) FindUnit(texture ITexture) (*Unit, error) {
	man.populate()

	for _, unit := range *man {
		if unit.Texture == nil || unit.Texture == texture {
			return unit, nil
		}
	}

	return nil, fmt.Errorf("no free texture unit found")
}