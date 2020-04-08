package Material

import "github.com/Adi146/goggle-engine/Core/Texture"

type IMaterial interface {
	Unbind()

	SetWrapMode(mode Texture.WrapMode)
	GenerateMibMaps(lodBias float32)
}
