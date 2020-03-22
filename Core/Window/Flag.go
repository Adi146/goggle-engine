package Window

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"gopkg.in/yaml.v3"
)

type Flag uint32

func (flag *Flag) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig []string
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	for _, flagString := range yamlConfig {
		switch flagString {
		case "fullscreen":
			*flag = *flag | sdl.WINDOW_FULLSCREEN
		case "fullscreen_desktop":
			*flag = *flag | sdl.WINDOW_FULLSCREEN_DESKTOP
		case "window_shown":
			*flag = *flag | sdl.WINDOW_SHOWN
		case "window_hidden":
			*flag = *flag | sdl.WINDOW_HIDDEN
		case "borderless":
			*flag = *flag | sdl.WINDOW_BORDERLESS
		case "minimized":
			*flag = *flag | sdl.WINDOW_MINIMIZED
		case "maximized":
			*flag = *flag | sdl.WINDOW_MAXIMIZED
		default:
			return fmt.Errorf("flag %s is not supported", yamlConfig)
		}
	}

	return nil
}
