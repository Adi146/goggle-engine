package Window_test

import (
	"github.com/Adi146/goggle-engine/Core/TestUtils"
	"testing"
)

func TestSDL(t *testing.T) {
	window, _ := TestUtils.CreateTestWindow(t)
	defer window.Destroy()
}
