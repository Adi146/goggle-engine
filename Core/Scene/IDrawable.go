package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type IDrawable interface {
	Draw(step *ProcessingPipelineStep) error
	GetPosition() *Vector.Vector3
}
