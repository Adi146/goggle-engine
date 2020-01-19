package Model

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	_ "image/png"
	"os"
)

type Texture struct {
	Source    string
	textureId uint32
}

func NewTextureFromFile(source string) (*Texture, error) {
	texture := Texture{
		Source: source,
	}

	img, err := DecodeImage(source)
	if err != nil {
		return nil, err
	}

	gl.GenTextures(1, &texture.textureId)
	gl.BindTexture(gl.TEXTURE_2D, texture.textureId)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return &texture, nil
}

func (tex *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tex.textureId)
}

func (tex *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func DecodeImage(imgfile string) (*image.RGBA, error) {
	file, err := os.Open(imgfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			rgba.Set(x, bounds.Dy()-y, img.At(x, y))
		}
	}

	return rgba, nil
}
