package Window

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"gopkg.in/yaml.v3"
)

const (
	yaml_key_fullscreen = "fullscreen"
	yaml_key_desktop    = "fullscreen_desktop"
	yaml_key_shown      = "window_shown"
	yaml_key_hidden     = "window_hidden"
	yaml_key_borderless = "borderless"
	yaml_key_minimized  = "minimized"
	yaml_key_maximized  = "maximized"
)

type Flag uint32

func (flag *Flag) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig []string
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	for _, flagString := range yamlConfig {

		switch flagString {
		case yaml_key_fullscreen:
			*flag = *flag | sdl.WINDOW_FULLSCREEN
		case yaml_key_desktop:
			*flag = *flag | sdl.WINDOW_FULLSCREEN_DESKTOP
		case yaml_key_shown:
			*flag = *flag | sdl.WINDOW_SHOWN
		case yaml_key_hidden:
			*flag = *flag | sdl.WINDOW_HIDDEN
		case yaml_key_borderless:
			*flag = *flag | sdl.WINDOW_BORDERLESS
		case yaml_key_minimized:
			*flag = *flag | sdl.WINDOW_MINIMIZED
		case yaml_key_maximized:
			*flag = *flag | sdl.WINDOW_MAXIMIZED
		default:
			return fmt.Errorf("flag %s is not supported", yamlConfig)
		}
	}

	return nil
}
