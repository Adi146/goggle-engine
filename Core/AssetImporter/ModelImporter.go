package AssetImporter

import (
	"fmt"
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	"path"
	"strings"
)

func ImportModel(filename string) (*Model.Model, *ImportResult) {
	result := newImportResult()

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

	materials := make([]*Model.Material, assimpScene.NumMaterials())
	meshes := make([]Model.MeshesWithMaterial, assimpScene.NumMeshes())

	for i, assimpMaterial := range assimpScene.Materials() {
		material, materialResult := importAssimpMaterial(assimpMaterial, path.Dir(filename))
		result.addResult(materialResult)
		materials[i] = material
	}

	for i, assimpMesh := range assimpScene.Meshes() {
		mesh, meshResult := importAssimpMesh(assimpMesh)
		result.addResult(meshResult)
		if !result.Success() {
			continue
		}
		meshes[i].Mesh = mesh
		meshes[i].Material = materials[assimpMesh.MaterialIndex()]
	}

	return &Model.Model{
		Meshes:      meshes,
		ModelMatrix: Matrix.Identity(),
	}, result
}

func importAssimpMaterial(assimpMaterial *assimp.Material, modelDir string) (*Model.Material, *ImportResult) {
	result := newImportResult()

	assimpDiffuse, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorDiffuse, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.addWarning(fmt.Errorf("could not load diffuse color"))
	}
	assimpSpecular, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorSpecular, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.addWarning(fmt.Errorf("could not load specular color"))
	}
	assimpEmissive, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorEmissive, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.addWarning(fmt.Errorf("could not load emissive color"))
	}
	assimpShininess, returnCode := assimpMaterial.GetMaterialFloat(assimp.MatKey_Shininess, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.addWarning(fmt.Errorf("could not load shininess"))
	}
	diffuseTextures, diffuseResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Diffuse, modelDir)
	result.addResult(diffuseResult)

	normalsTexture, normalsResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Normals, modelDir)
	result.addResult(normalsResult)

	return &Model.Material{
		DiffuseBaseColor:  Vector.Vector3{assimpDiffuse.R(), assimpDiffuse.G(), assimpDiffuse.B()},
		SpecularBaseColor: Vector.Vector3{assimpSpecular.R(), assimpSpecular.G(), assimpSpecular.B()},
		EmissiveBaseColor: Vector.Vector3{assimpEmissive.R(), assimpEmissive.G(), assimpEmissive.B()},

		Shininess: assimpShininess,

		DiffuseTextures: diffuseTextures,
		NormalTextures:  normalsTexture,
	}, result
}

func importTexturesOfAssimpMaterial(assimpMaterial *assimp.Material, textureType assimp.TextureMapping, modelDir string) ([]*Model.Texture, *ImportResult) {
	result := newImportResult()
	var textures []*Model.Texture

	numTextures := assimpMaterial.GetMaterialTextureCount(assimp.TextureType(textureType))
	for i := 0; i < numTextures; i++ {
		textureFile, mapping, uvIndex, blend, op, mapmode, flags, returnCode := assimpMaterial.GetMaterialTexture(assimp.TextureType(textureType), i)
		if returnCode != assimp.Return_Success {
			result.addWarning(fmt.Errorf("could not get texture for material with index %d", i))
			continue
		}
		fmt.Println(textureFile, mapping, uvIndex, blend, op, mapmode, flags)

		if strings.HasSuffix("*/", textureFile) {
			result.addWarning(fmt.Errorf("embedded textures are not supported yet"))
			continue
		} else {
			texture, textureResult := ImportTexture(path.Join(modelDir, textureFile))
			result.addResultAsWarning(textureResult)
			if !textureResult.Success() {
				continue
			}
			textures = append(textures, texture)
		}
	}

	return textures, result
}

func importAssimpMesh(assimpMesh *assimp.Mesh) (*Model.Mesh, *ImportResult) {
	result := newImportResult()

	assimpVertices := assimpMesh.Vertices()
	assimpNormals := assimpMesh.Normals()
	assimpUVs := assimpMesh.TextureCoords(0)
	assimpFaces := assimpMesh.Faces()
	assimpTangents := assimpMesh.Tangents()

	vertices := make([]Buffer.Vertex, assimpMesh.NumVertices())
	for i := 0; i < assimpMesh.NumVertices(); i++ {
		vertices[i].Position = Vector.Vector3{assimpVertices[i].X(), assimpVertices[i].Y(), assimpVertices[i].Z()}
		vertices[i].Normal = Vector.Vector3{assimpNormals[i].X(), assimpNormals[i].Y(), assimpNormals[i].Z()}
		vertices[i].UV = Vector.Vector2{assimpUVs[i].X(), assimpUVs[i].Y()}
		vertices[i].Tangent = Vector.Vector3{assimpTangents[i].X(), assimpTangents[i].Y(), assimpTangents[i].Z()}
	}

	var indices []uint32
	for _, assimpFace := range assimpFaces {
		indices = append(indices, assimpFace.CopyIndices()...)
	}

	mesh, err := Model.NewMesh(vertices, Buffer.RegisterVertexBufferAttributes, indices)
	if err != nil {
		result.addError(err)
		return nil, result
	}

	result.NumImportedAssets = 1
	return mesh, result
}
