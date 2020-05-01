package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type InstancedMesh struct {
	ModelMatrix               GeometryMath.Matrix4x4
	boundingVolume            BoundingVolume.IBoundingVolume
	TransformedBoundingVolume BoundingVolume.IBoundingVolume
	FrustumCulling            bool
	Master                    *InstanceMasterMesh
}

func (mesh *InstancedMesh) GetModelMatrix() GeometryMath.Matrix4x4 {
	return mesh.ModelMatrix
}

func (mesh *InstancedMesh) SetModelMatrix(mat GeometryMath.Matrix4x4) {
	mesh.ModelMatrix = mat
	mesh.TransformedBoundingVolume = mesh.boundingVolume.Transform(mat.Mul(mesh.Master.MasterMatrix))
}

func (mesh *InstancedMesh) GetBoundingVolume() BoundingVolume.IBoundingVolume {
	return mesh.TransformedBoundingVolume
}

func (mesh *InstancedMesh) EnableFrustumCulling() {
	mesh.FrustumCulling = true
}

func (mesh *InstancedMesh) DisableFrustumCulling() {
	mesh.FrustumCulling = false
}
