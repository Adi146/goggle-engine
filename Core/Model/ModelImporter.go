package Model

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"path"
	"strings"

	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

var textureTypeMap = map[assimp.TextureMapping]Texture.Type{
	assimp.TextureMapping_Diffuse:  Texture.DiffuseTexture,
	assimp.TextureMapping_Specular: Texture.SpecularTexture,
	assimp.TextureMapping_Emissive: Texture.EmissiveTexture,
	assimp.TextureMapping_Normals:  Texture.NormalsTexture,
}

func ImportModel(filename string, index int) (*Model, ImportResult) {
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

	materials := make([]*Material.Material, assimpScene.NumMaterials())
	models := make([]Model, assimpScene.NumMeshes())

	for i, assimpMaterial := range assimpScene.Materials() {
		material, materialResult := importAssimpMaterial(assimpMaterial, path.Dir(filename))
		result.Errors.Push(&materialResult.Errors)
		result.Warnings.Push(&materialResult.Warnings)
		materials[i] = material
	}

	for i, assimpMesh := range assimpScene.Meshes() {
		models[i].IMesh = importAssimpMesh(assimpMesh)
		models[i].Material = materials[assimpMesh.MaterialIndex()]
	}

	return &models[index], result
}

func importAssimpMaterial(assimpMaterial *assimp.Material, modelDir string) (*Material.Material, ImportResult) {
	var result ImportResult
	material := Material.Material{
		UvScale: 1,
	}

	diffuseColor, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorDiffuse, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load diffuse color"))
	} else {
		material.DiffuseBaseColor = GeometryMath.Vector4{diffuseColor.R(), diffuseColor.G(), diffuseColor.B(), diffuseColor.A()}
	}
	specularColor, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorSpecular, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load specular color"))
	} else {
		material.SpecularBaseColor = GeometryMath.Vector3{specularColor.R(), specularColor.G(), specularColor.B()}
	}
	emissiveColor, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorEmissive, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load emissive color"))
	} else {
		material.EmissiveBaseColor = GeometryMath.Vector3{emissiveColor.R(), emissiveColor.G(), emissiveColor.B()}
	}
	shininess, returnCode := assimpMaterial.GetMaterialFloat(assimp.MatKey_Shininess, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		result.Warnings.Push(fmt.Errorf("could not load shininess"))
	} else {
		material.Shininess = shininess
	}

	diffuseTextures, textureResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Diffuse, modelDir)
	result.Errors.Push(&textureResult.Errors)
	result.Warnings.Push(&textureResult.Warnings)
	if textureResult.Success() && len(diffuseTextures) > 0 {
		material.Textures.Diffuse = diffuseTextures[0]
	}

	specularTextures, textureResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Specular, modelDir)
	result.Errors.Push(&textureResult.Errors)
	result.Warnings.Push(&textureResult.Warnings)
	if textureResult.Success() && len(specularTextures) > 0 {
		material.Textures.Specular = specularTextures[0]
	}

	emissiveTextures, textureResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Emissive, modelDir)
	result.Errors.Push(&textureResult.Errors)
	result.Warnings.Push(&textureResult.Warnings)
	if textureResult.Success() && len(emissiveTextures) > 0 {
		material.Textures.Emissive = emissiveTextures[0]
	}

	normalTextures, textureResult := importTexturesOfAssimpMaterial(assimpMaterial, assimp.TextureMapping_Normals, modelDir)
	result.Errors.Push(&textureResult.Errors)
	result.Warnings.Push(&textureResult.Warnings)
	if textureResult.Success() && len(normalTextures) > 0 {
		material.Textures.Normal = normalTextures[0]
	}

	return &material, result
}

func importTexturesOfAssimpMaterial(assimpMaterial *assimp.Material, textureType assimp.TextureMapping, modelDir string) ([]*Texture.Texture2D, ImportResult) {
	var result ImportResult
	var textures []*Texture.Texture2D

	numTextures := assimpMaterial.GetMaterialTextureCount(assimp.TextureType(textureType))
	for i := 0; i < numTextures; i++ {
		textureFile, _, _, _, _, _, _, returnCode := assimpMaterial.GetMaterialTexture(assimp.TextureType(textureType), i)
		if returnCode != assimp.Return_Success {
			result.Warnings.Push(fmt.Errorf("could not get texture for material with index %d", i))
			continue
		}

		if strings.HasSuffix("*/", textureFile) {
			result.Warnings.Push(fmt.Errorf("embedded textures are not supported yet"))
			continue
		} else {
			texture, err := Texture.ImportTexture(path.Join(modelDir, textureFile), textureTypeMap[textureType])
			if err != nil {
				result.Warnings.Push(err)
				continue
			}
			textures = append(textures, texture)
		}
	}

	return textures, result
}

func importAssimpMesh(assimpMesh *assimp.Mesh) Mesh.IMesh {
	assimpVertices := assimpMesh.Vertices()
	assimpNormals := assimpMesh.Normals()
	assimpUVs := assimpMesh.TextureCoords(0)
	assimpFaces := assimpMesh.Faces()
	assimpTangents := assimpMesh.Tangents()
	assimpBiTangents := assimpMesh.Bitangents()

	vertices := make([]Mesh.Vertex, assimpMesh.NumVertices())
	for i := 0; i < assimpMesh.NumVertices(); i++ {
		vertices[i] = Mesh.Vertex{
			Position:  GeometryMath.Vector3{assimpVertices[i].X(), assimpVertices[i].Y(), assimpVertices[i].Z()},
			UV:        GeometryMath.Vector2{assimpUVs[i].X(), assimpUVs[i].Y()},
			Normal:    GeometryMath.Vector3{assimpNormals[i].X(), assimpNormals[i].Y(), assimpNormals[i].Z()},
			Tangent:   GeometryMath.Vector3{assimpTangents[i].X(), assimpTangents[i].Y(), assimpTangents[i].Z()},
			BiTangent: GeometryMath.Vector3{assimpBiTangents[i].X(), assimpBiTangents[i].Y(), assimpBiTangents[i].Z()},
		}
	}

	var indices []uint32
	for _, assimpFace := range assimpFaces {
		indices = append(indices, assimpFace.CopyIndices()...)
	}

	return Mesh.NewMesh(vertices, indices)
}
