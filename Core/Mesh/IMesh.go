package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type IMesh interface {
	BoundingVolume.ICollisionObject

	GetModelMatrix() GeometryMath.Matrix4x4
	SetModelMatrix(mat GeometryMath.Matrix4x4)

	EnableFrustumCulling()
	DisableFrustumCulling()
}
