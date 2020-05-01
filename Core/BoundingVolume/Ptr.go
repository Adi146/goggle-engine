package BoundingVolume

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

const (
	Yaml_key_aabb   = "AABB"
	Yaml_key_sphere = "sphere"
	Yaml_key_point  = "point"
	Yaml_key_empty  = ""
)

type Ptr struct {
	IBoundingVolume
}

func (ptr *Ptr) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		Type string `yaml:"type"`
	}

	var err error
	var configAvailable bool
	if value.Kind == yaml.ScalarNode {
		err = value.Decode(&yamlConfig.Type)
		configAvailable = false
	} else {
		err = value.Decode(&yamlConfig)
		configAvailable = true
	}

	if err != nil {
		return err
	}

	switch yamlConfig.Type {
	case Yaml_key_aabb:
		ptr.IBoundingVolume = NewDefaultAABB()
	case Yaml_key_sphere:
		ptr.IBoundingVolume = NewDefaultSphere()
	case Yaml_key_point:
		ptr.IBoundingVolume = NewDefaultPoint()
	case Yaml_key_empty:
		if ptr.IBoundingVolume == nil {
			return fmt.Errorf("no bounding volume type specified")
		}
	default:
		return fmt.Errorf("%s is not a bounding volume type", yamlConfig.Type)
	}

	if configAvailable {
		return value.Decode(ptr.IBoundingVolume)
	}

	return nil
}
