package Window

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type SyncInterval int

func (sync *SyncInterval) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig string
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	switch yamlConfig {
	case "normal", "":
		*sync = 0
	case "adaptive":
		*sync = -1
	case "vertical":
		*sync = 1
	default:
		return fmt.Errorf("sync %s not supported", yamlConfig)
	}

	return nil
}
