package MatrixFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"gopkg.in/yaml.v3"
	"reflect"
)

var (
	typeFactory = map[string]reflect.Type{}
)

func AddType(key string, configType reflect.Type) {
	typeFactory[key] = configType
}

func Get(key string, yaml yaml.Node) (*Matrix.Matrix4x4, error) {
	matrixType, ok := typeFactory[key]
	if !ok {
		return nil, fmt.Errorf("matrix type %s not in factory", key)
	}

	matrixConfig := reflect.New(matrixType).Interface().(IYamlMatrixConfig)
	yaml.Decode(matrixConfig)

	return matrixConfig.Decode(), nil
}

type IYamlMatrixConfig interface {
	Decode() *Matrix.Matrix4x4
}
