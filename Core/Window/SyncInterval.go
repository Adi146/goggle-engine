package Window

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Utils/Constants"
	"gopkg.in/yaml.v3"
)

const (
	yaml_key_normal   = "normal"
	yaml_key_adaptive = "adaptive"
	yaml_key_vertical = "vertical"
)

type SyncInterval int

func (sync *SyncInterval) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig string
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	switch yamlConfig {
	case yaml_key_normal, Constants.Yaml_key_false, "":
		*sync = 0
	case yaml_key_adaptive:
		*sync = -1
	case yaml_key_vertical, Constants.Yaml_key_true:
		*sync = 1
	default:
		return fmt.Errorf("sync %s not supported", yamlConfig)
	}

	return nil
}
