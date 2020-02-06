package AssetImporter

import (
	"fmt"
	"path"
	"strings"

	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

var textureTypeMap = map[assimp.TextureMapping]Texture.TextureType{
	assimp.TextureMapping_Diffuse:  Texture.DiffuseTexture,
	assimp.TextureMapping_Specular: Texture.SpecularTexture,
	assimp.TextureMapping_Emissive: Texture.EmissiveTexture,
	assimp.TextureMapping_Normals:  Texture.NormalsTexture,
}

func ImportModel(filename string) (*Model.Model, ImportResult) {
	var result ImportResult

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
		result.Errors.Push(&materialResult.Errors)
		result.Warnings.Push(&materialResult.Warnings)
		materials[i] = material
	}

	for i, assimpMesh := range assimpScene.Meshes() {
		mesh, meshResult := importAssimpMesh(assimpMesh)
		result.Errors.Push(&meshResult.Errors)
		result.Warnings.Push(&meshResult.Warnings)
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

func importAssimpMaterial(assimpMaterial *assimp.Material, modelDir string) (*Model.Material, ImportResult) {
	var result ImportResult

	assimpDiffuse, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorDiffuse, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load diffuse color"))
	}
	assimpSpecular, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorSpecular, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load specular color"))
	}
	assimpEmissive, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorEmissive, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load emissive color"))
	}
	assimpShininess, returnCode := assimpMaterial.GetMaterialFloat(assimp.MatKey_Shininess, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load shininess"))
	}

	var modelTextures []*Texture.Texture
	for textureType, _ := range textureTypeMap {
		textures, textureResult := importTexturesOfAssimpMaterial(assimpMaterial, textureType, modelDir)
		result.Errors.Push(&textureResult.Errors)
		result.Warnings.Push(&textureResult.Warnings)

		modelTextures = append(modelTextures, textures...)
	}

	return &Model.Material{
		DiffuseBaseColor:  Vector.Vector3{assimpDiffuse.R(), assimpDiffuse.G(), assimpDiffuse.B()},
		SpecularBaseColor: Vector.Vector3{assimpSpecular.R(), assimpSpecular.G(), assimpSpecular.B()},
		EmissiveBaseColor: Vector.Vector3{assimpEmissive.R(), assimpEmissive.G(), assimpEmissive.B()},

		Shininess: assimpShininess,

		Textures: modelTextures,
	}, result
}

func importTexturesOfAssimpMaterial(assimpMaterial *assimp.Material, textureType assimp.TextureMapping, modelDir string) ([]*Texture.Texture, ImportResult) {
	var result ImportResult
	var textures []*Texture.Texture

	numTextures := assimpMaterial.GetMaterialTextureCount(assimp.TextureType(textureType))
	for i := 0; i < numTextures; i++ {
		textureFile, mapping, uvIndex, blend, op, mapmode, flags, returnCode := assimpMaterial.GetMaterialTexture(assimp.TextureType(textureType), i)
		if returnCode != assimp.Return_Success {
			result.Warnings.Push(fmt.Errorf("could not get texture for material with index %d", i))
			continue
		}
		fmt.Println(textureFile, mapping, uvIndex, blend, op, mapmode, flags)

		if strings.HasSuffix("*/", textureFile) {
			result.Warnings.Push(fmt.Errorf("embedded textures are not supported yet"))
			continue
		} else {
			texture, textureResult := ImportTexture(path.Join(modelDir, textureFile), textureTypeMap[textureType])
			result.Warnings.Push(&textureResult.Warnings)
			result.Warnings.Push(&textureResult.Errors)
			if !textureResult.Success() {
				continue
			}
			textures = append(textures, texture)
		}
	}

	return textures, result
}

func importAssimpMesh(assimpMesh *assimp.Mesh) (*Model.Mesh, ImportResult) {
	var result ImportResult

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
	result.Errors.Push(err)

	return mesh, result
}
