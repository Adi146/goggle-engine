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
	Decode() *Matrix4x4
}

func (m1 *Matrix4x4) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.SequenceNode {
		var yamlConfig []Matrix4x4
		if err := value.Decode(&yamlConfig); err != nil {
			return err
		}

		*m1 = *Identity()
		for _, matrix := range yamlConfig {
			*m1 = *m1.Mul(&matrix)
		}
	} else {
		var yamlConfig map[string]yaml.Node
		if err := value.Decode(&yamlConfig); err != nil {
			return err
		}

		*m1 = *Identity()
		for yamlKey, tmpValue := range yamlConfig {
			config, err := getMatrixConfig(yamlKey)
			if err != nil {
				return err
			}

			if err := tmpValue.Decode(config); err != nil {
				return err
			}
			*m1 = *m1.Mul(config.Decode())
		}
	}

	return nil
}

func getMatrixConfig(yamlKey string) (IMatrixConfig, error) {
	switch yamlKey {
	case yaml_key_orthogonal:
		return new(orthogonalConfig), nil
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

type orthogonalConfig struct {
	Left   float32 `yaml:"left"`
	Right  float32 `yaml:"right"`
	Bottom float32 `yaml:"bottom"`
	Top    float32 `yaml:"top"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (config *orthogonalConfig) Decode() *Matrix4x4 {
	return Orthographic(
		config.Left,
		config.Right,
		config.Bottom,
		config.Top,
		config.Near,
		config.Far,
	)
}

type PerspectiveConfig struct {
	Fovy   float32 `yaml:"fovy"`
	Aspect float32 `yaml:"aspect"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (config *PerspectiveConfig) Decode() *Matrix4x4 {
	return Perspective(
		Radians(config.Fovy),
		config.Aspect,
		config.Near,
		config.Far,
	)
}

type rotationConfig struct {
	Vector Vector3 `yaml:"axis"`
	Angle  float32 `yaml:"angle"`
}

func (config *rotationConfig) Decode() *Matrix4x4 {
	return Rotate(Radians(config.Angle), &config.Vector)
}

type scaleConfig float32

func (config *scaleConfig) Decode() *Matrix4x4 {
	return Scale(float32(*config))
}

type translationConfig Vector3

func (config *translationConfig) Decode() *Matrix4x4 {
	return Translate((*Vector3)(config))
}
