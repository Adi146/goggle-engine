package Model

import (
	"encoding/binary"
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
