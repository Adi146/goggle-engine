package Function

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Utils/Constants"
	"gopkg.in/yaml.v3"

	"github.com/go-gl/gl/v4.3-core/gl"
)

const (
	yaml_key_front        = "front"
	yaml_key_back         = "back"
	yaml_key_frontAndBack = "frontAndBack"
)

var (
	Front = CullFunction{
		Enabled:  true,
		Function: gl.FRONT,
	}
	Back = CullFunction{
		Enabled:  true,
		Function: gl.BACK,
	}
	FrontAndBack = CullFunction{
		Enabled:  true,
		Function: gl.FRONT_AND_BACK,
	}
	CullDisabled = CullFunction{
		Enabled: false,
	}
)

type CullFunction struct {
	Enabled  bool
	Function uint32
}

func (f *CullFunction) Set() {
	if f.Enabled {
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(f.Function)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
}

func (f *CullFunction) UnmarshalYAML(value *yaml.Node) error {
	var tmp string
	value.Decode(&tmp)

	switch tmp {
	case yaml_key_front:
		*f = Front
	case yaml_key_back, Constants.Yaml_key_true:
		*f = Back
	case yaml_key_frontAndBack:
		*f = FrontAndBack
	case Constants.Yaml_key_false:
		*f = CullDisabled
	default:
		return fmt.Errorf("%s is not a cull function", tmp)
	}

	return nil
}

func GetCurrentCullFunction() *CullFunction {
	var enabled bool
	var function int32

	gl.GetBooleanv(gl.CULL_FACE, &enabled)
	gl.GetIntegerv(gl.CULL_FACE_MODE, &function)

	return &CullFunction{
		Enabled:  enabled,
		Function: uint32(function),
	}
}
