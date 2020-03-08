package FrameBuffer

import (
	"fmt"
	"gopkg.in/yaml.v3"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type CullFunction struct {
	Enabled  bool
	Function uint32
}

func (f *CullFunction) UnmarshalYAML(value *yaml.Node) error {
	var tmp string
	value.Decode(&tmp)

	switch tmp {
	case "front":
		f.Enabled = true
		f.Function = gl.FRONT
	case "back", "true":
		f.Enabled = true
		f.Function = gl.BACK
	case "front_and_back":
		f.Enabled = true
		f.Function = gl.FRONT_AND_BACK
	case "false":
		f.Enabled = false
	default:
		return fmt.Errorf("%s is not a cull function", tmp)
	}

	return nil
}
