package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"gopkg.in/yaml.v3"
)

type MatrixConfig struct {
	Matrix.Matrix4x4
}

type tmpConfig map[string]yaml.Node

func (config *MatrixConfig) UnmarshalYAML(value *yaml.Node) error {
	var tmpConfig tmpConfig
	value.Decode(&tmpConfig)

	tmpMatrix := Matrix.Identity()
	for factoryKey, tmpValue := range tmpConfig {
		mat, err := Get(factoryKey, tmpValue)
		if err != nil {
			return err
		}

		tmpMatrix = tmpMatrix.Mul(mat)
	}

	config.Matrix4x4 = *tmpMatrix
	return nil
}
