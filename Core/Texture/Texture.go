package Texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Type string
type WrapMode int32

const (
	Repeat         WrapMode = gl.REPEAT
	MirroredRepeat WrapMode = gl.MIRRORED_REPEAT
	ClampToEdge    WrapMode = gl.CLAMP_TO_EDGE
	ClampToBorder  WrapMode = gl.CLAMP_TO_BORDER
)

type Texture struct {
	ID     uint32
	Target uint32
	Type   Type
	Unit   *Unit
}

func (texture *Texture) GetID() uint32 {
	return texture.ID
}

func (texture *Texture) GetUnit() *Unit {
	return texture.Unit
}

func (texture *Texture) GetType() Type {
	return texture.Type
}

func (texture *Texture) GetTarget() uint32 {
	return texture.Target
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

func (texture *Texture) SetWrapMode(mode WrapMode) {
	gl.TextureParameteri(texture.ID, gl.TEXTURE_WRAP_S, int32(mode))
	gl.TextureParameteri(texture.ID, gl.TEXTURE_WRAP_T, int32(mode))
	gl.TextureParameteri(texture.ID, gl.TEXTURE_WRAP_R, int32(mode))
}

func (texture *Texture) GenerateMipMap(lodBias float32) {
	gl.GenerateTextureMipmap(texture.ID)
	gl.TextureParameteri(texture.ID, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TextureParameterf(texture.ID, gl.TEXTURE_LOD_BIAS, lodBias)
}