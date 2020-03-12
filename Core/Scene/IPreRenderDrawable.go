package Scene

type IPreRenderObject interface {
	IDrawable
	GetPriority() int
}

type byPriority []IPreRenderObject

func (slice byPriority) Len() int {
	return len(slice)
}

func (slice byPriority) Less(i int, j int) bool {
	return slice[i].GetPriority() > slice[j].GetPriority()
}

func (slice byPriority) Swap(i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
