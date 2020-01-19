package Model

import (
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"path"
)

var textureTypes = [...]assimp.TextureMapping{
	assimp.TextureMapping_Diffuse,
	assimp.TextureMapping_Normals,
}

type MeshesWithMaterial struct {
	*Mesh
	*Material
}

type Model struct {
	Meshes      []MeshesWithMaterial
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
			assimp.Process_ImproveCacheLocality |
			assimp.Process_CalcTangentSpace,
	)

	materials := make([]*Material, assimpScene.NumMaterials())
	meshes := make([]MeshesWithMaterial, assimpScene.NumMeshes())

	for i, assimpMaterial := range assimpScene.Materials() {
		material, err := ImportMaterial(assimpMaterial, path.Dir(filename))
		if err != nil {
			return nil, err
		}
		materials[i] = material
	}

	for i, assimpMesh := range assimpScene.Meshes() {
		mesh, err := ImportMesh(assimpMesh)
		if err != nil {
			return nil, err
		}
		meshes[i].Mesh = mesh
		meshes[i].Material = materials[assimpMesh.MaterialIndex()]
	}

	return &Model{
		Meshes:      meshes,
		ModelMatrix: Matrix.Identity(),
	}, nil
}
