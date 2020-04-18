package Scene

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type ITransparentDrawable interface {
	IDrawable
	GetPosition() GeometryMath.Vector3
}

type transparentObject struct {
	IDrawable
	CameraDistance float32
}

type byDistance []transparentObject

func (slice byDistance) Len() int {
	return len(slice)
}

func (slice byDistance) Less(i int, j int) bool {
	return slice[i].CameraDistance > slice[j].CameraDistance
}

func (slice byDistance) Swap(i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
