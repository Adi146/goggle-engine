package Window

type IKeyboardInput interface {
	IsKeyPressed(key string) bool
}
