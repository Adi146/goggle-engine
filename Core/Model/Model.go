package Model

import (
	"encoding/binary"
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Geometry"
	"os"
)

type GeometryWithMaterial struct {
	*Geometry.Geometry
	*Material
}

type Model struct {
	Geometries []GeometryWithMaterial
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
		material, err := ImportMaterial(assimpMaterial)
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
		Geometries: geometries,
	}, nil
}

func NewModelFromFile(file *os.File) (*Model, error) {
	var numGeometries uint64
	if err := binary.Read(file, binary.LittleEndian, &numGeometries); err != nil {
		return nil, err
	}

	geometriesWithMaterial := make([]GeometryWithMaterial, numGeometries)

	for i := uint64(0); i < numGeometries; i++ {
		geometry, err := Geometry.NewGeometryFromFile(file)
		if err != nil {
			return nil, err
		}

		material, err := NewMaterialFromFile(file)
		if err != nil {
			return nil, err
		}

		geometriesWithMaterial[i] = GeometryWithMaterial{
			Geometry: geometry,
			Material: material,
		}
	}

	return &Model{
		Geometries: geometriesWithMaterial,
	}, nil
}
