package Model

import (
	"encoding/binary"
	"fmt"
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"os"
	"path"
	"strings"
)

type Material struct {
	DiffuseBaseColor  Vector.Vector3
	SpecularBaseColor Vector.Vector3
	EmissiveBaseColor Vector.Vector3

	Shininess float32

	DiffuseTextures []*Texture.Texture
	NormalTextures  []*Texture.Texture
}

func NewMaterialFromFile(file *os.File) (*Material, error) {
	material := Material{}

	if err := binary.Read(file, binary.LittleEndian, &material.DiffuseBaseColor); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.SpecularBaseColor); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.EmissiveBaseColor); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.Shininess); err != nil {
		return nil, err
	}

	return &material, nil
}

func ImportMaterial(assimpMaterial *assimp.Material, modelDir string) (*Material, error) {
	assimpDiffuse, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorDiffuse, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		return nil, fmt.Errorf("could not load diffuse color from material")
	}
	assimpSpecular, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorSpecular, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		return nil, fmt.Errorf("could not load specular color from material")
	}
	assimpEmissive, returnCode := assimpMaterial.GetMaterialColor(assimp.MatKey_ColorEmissive, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		return nil, fmt.Errorf("could not load emissive color from material")
	}
	assimpShininess, returnCode := assimpMaterial.GetMaterialFloat(assimp.MatKey_Shininess, assimp.TextureType(assimp.TextureMapping_None), 0)
	if returnCode != assimp.Return_Success {
		return nil, fmt.Errorf("could not load shininess from material")
	}
	diffuseTextures, err := ImportTexturesOfMaterial(assimpMaterial, assimp.TextureMapping_Diffuse, modelDir)
	if err != nil {
		return nil, err
	}
	normalsTexture, err := ImportTexturesOfMaterial(assimpMaterial, assimp.TextureMapping_Normals, modelDir)
	if err != nil {
		return nil, err
	}

	return &Material{
		DiffuseBaseColor:  Vector.Vector3{assimpDiffuse.R(), assimpDiffuse.G(), assimpDiffuse.B()},
		SpecularBaseColor: Vector.Vector3{assimpSpecular.R(), assimpSpecular.G(), assimpSpecular.B()},
		EmissiveBaseColor: Vector.Vector3{assimpEmissive.R(), assimpEmissive.G(), assimpEmissive.B()},

		Shininess: assimpShininess,

		DiffuseTextures: diffuseTextures,
		NormalTextures:  normalsTexture,
	}, nil
}

func ImportTexturesOfMaterial(assimpMaterial *assimp.Material, textureType assimp.TextureMapping, modelDir string) ([]*Texture.Texture, error) {
	var textures []*Texture.Texture
	var textureFiles []string

	numTextures := assimpMaterial.GetMaterialTextureCount(assimp.TextureType(textureType))
	for i := 0; i < numTextures; i++ {
		textureFile, mapping, uvIndex, blend, op, mapmode, flags, returnCode := assimpMaterial.GetMaterialTexture(assimp.TextureType(textureType), i)
		if returnCode != assimp.Return_Success {
			return nil, fmt.Errorf("could not get texture for material with index %d", i)
		}
		fmt.Println(textureFile, mapping, uvIndex, blend, op, mapmode, flags)

		if strings.HasSuffix("*/", textureFile) {
			return nil, fmt.Errorf("embedded textures are not supported yet")
		}

		textureFiles = append(textureFiles, path.Join(modelDir, textureFile))
	}

	for _, textureFile := range textureFiles {
		texture, err := Texture.NewTextureFromFile(textureFile)
		if err != nil {
			return nil, err
		}
		textures = append(textures, texture)
	}

	return textures, nil
}
