package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IDrawable interface {
	Draw(step *ProcessingPipelineStep) error
	GetPosition() *GeometryMath.Vector3
}
