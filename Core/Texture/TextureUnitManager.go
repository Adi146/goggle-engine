package Texture

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type assignedUnit struct {
	texture  ITexture
	reserved bool
}

var textureUnits []*assignedUnit

func FindFreeUnit(texture ITexture) (uint32, error) {
	if len(textureUnits) == 0 {
		var tmp int32
		gl.GetIntegerv(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS, &tmp)
		textureUnits = make([]*assignedUnit, tmp)
	}

	for unit, assignedTexture := range textureUnits {
		if unit != 0 {
			if assignedTexture == nil {
				return uint32(unit), nil
			} else if assignedTexture.texture == texture {
				return uint32(unit), nil
			}
		}
	}

	return 0, fmt.Errorf("no free texture unit found")
}

func BindTexture(texture ITexture, unit uint32, reserve bool) {
	texture.Bind(unit)

	textureUnits[unit] = &assignedUnit{
		texture:  texture,
		reserved: reserve,
	}
}

func Clear() {
	for unit, assignedTexture := range textureUnits {
		if assignedTexture != nil && !assignedTexture.reserved {
			textureUnits[unit] = nil
		}
	}
}
