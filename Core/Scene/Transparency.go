package Scene

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
