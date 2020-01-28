package Factory

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
)

type Config struct {
	Scene        *Scene.Scene
	FrameBuffers []FrameBuffer.IFrameBuffer
}
