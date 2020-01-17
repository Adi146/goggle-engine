package Model

import (
	"encoding/binary"
	"fmt"
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"os"
)

type Material struct {
	Diffuse   Vector.Vector3
	Specular  Vector.Vector3
	Emissive  Vector.Vector3
	Shininess float32
}

func NewMaterialFromFile(file *os.File) (*Material, error) {
	material := Material{}

	if err := binary.Read(file, binary.LittleEndian, &material.Diffuse); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.Specular); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.Emissive); err != nil {
		return nil, err
	}
	if err := binary.Read(file, binary.LittleEndian, &material.Shininess); err != nil {
		return nil, err
	}

	return &material, nil
}

func ImportMaterial(assimpMaterial *assimp.Material) (*Material, error) {
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

	return &Material{
		Diffuse:   Vector.Vector3{assimpDiffuse.R(), assimpDiffuse.G(), assimpDiffuse.B()},
		Specular:  Vector.Vector3{assimpSpecular.R(), assimpSpecular.G(), assimpSpecular.B()},
		Emissive:  Vector.Vector3{assimpEmissive.R(), assimpEmissive.G(), assimpEmissive.B()},
		Shininess: assimpShininess,
	}, nil
}
