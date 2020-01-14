package Model

import (
	"encoding/binary"
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
