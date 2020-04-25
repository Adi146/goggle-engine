package GeometryMath

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

const (
	yaml_key_orthogonal  = "orthogonal"
	yaml_key_perspective = "perspective"
	yaml_key_rotation    = "rotation"
	yaml_key_scale       = "scale"
	yaml_key_translation = "translation"
)

type IMatrixConfig interface {
	Decode() Matrix4x4
}

func (m1 *Matrix4x4) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.SequenceNode {
		var yamlConfig []Matrix4x4
		if err := value.Decode(&yamlConfig); err != nil {
			return err
		}

		*m1 = Identity()
		for _, matrix := range yamlConfig {
			*m1 = m1.Mul(matrix)
		}
	} else {
		var yamlConfig map[string]yaml.Node
		if err := value.Decode(&yamlConfig); err != nil {
			return err
		}

		*m1 = Identity()
		for yamlKey, tmpValue := range yamlConfig {
			config, err := getMatrixConfig(yamlKey)
			if err != nil {
				return err
			}

			if err := tmpValue.Decode(config); err != nil {
				return err
			}
			*m1 = m1.Mul(config.Decode())
		}
	}

	return nil
}

func getMatrixConfig(yamlKey string) (IMatrixConfig, error) {
	switch yamlKey {
	case yaml_key_orthogonal:
		return new(OrthographicConfig), nil
	case yaml_key_perspective:
		return new(PerspectiveConfig), nil
	case yaml_key_rotation:
		return new(rotationConfig), nil
	case yaml_key_scale:
		return new(scaleConfig), nil
	case yaml_key_translation:
		return new(translationConfig), nil
	default:
		return nil, fmt.Errorf("matrix key %s is not supported", yamlKey)
	}
}

type rotationConfig struct {
	Vector Vector3 `yaml:"axis"`
	Angle  float32 `yaml:"angle"`
}

func (config *rotationConfig) Decode() Matrix4x4 {
	return Rotate(Radians(config.Angle), config.Vector)
}

type scaleConfig float32

func (config *scaleConfig) Decode() Matrix4x4 {
	return Scale(float32(*config))
}

type translationConfig Vector3

func (config *translationConfig) Decode() Matrix4x4 {
	return Translate((Vector3)(*config))
}
