package Model

import (
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Geometry"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"path"
)

var textureTypes = [...]assimp.TextureMapping{
	assimp.TextureMapping_Diffuse,
	assimp.TextureMapping_Normals,
}

type GeometryWithMaterial struct {
	*Geometry.Geometry
	*Material
}

type Model struct {
	Geometries  []GeometryWithMaterial
	ModelMatrix *Matrix.Matrix4x4
}

func ImportModel(filename string) (*Model, error) {
	assimpScene := assimp.ImportFile(filename, 0)
	assimpScene.ApplyPostProcessing(
		assimp.Process_PreTransformVertices |
			assimp.Process_Triangulate |
			assimp.Process_GenNormals |
			assimp.Process_OptimizeMeshes |
			assimp.Process_OptimizeGraph |
			assimp.Process_JoinIdenticalVertices |
			assimp.Process_ImproveCacheLocality,
	)

	materials := make([]*Material, assimpScene.NumMaterials())
	geometries := make([]GeometryWithMaterial, assimpScene.NumMeshes())

	for i, assimpMaterial := range assimpScene.Materials() {
		material, err := ImportMaterial(assimpMaterial, path.Dir(filename))
		if err != nil {
			return nil, err
		}
		materials[i] = material
	}

	for i, assimpMesh := range assimpScene.Meshes() {
		geometry, err := Geometry.ImportGeometry(assimpMesh)
		if err != nil {
			return nil, err
		}
		geometries[i].Geometry = geometry
		geometries[i].Material = materials[assimpMesh.MaterialIndex()]
	}

	return &Model{
		Geometries:  geometries,
		ModelMatrix: Matrix.Identity(),
	}, nil
}
