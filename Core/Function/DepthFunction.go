package Function

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Utils/Constants"
	"github.com/go-gl/gl/v4.3-core/gl"
	"gopkg.in/yaml.v3"
)

const (
	yaml_key_never        = "never"
	yaml_key_less         = "less"
	yaml_key_equal        = "equal"
	yaml_key_lessEqual    = "lessEqual"
	yaml_key_greater      = "greater"
	yaml_key_notEqual     = "notEqual"
	yaml_key_greaterEqual = "greaterEqual"
	yaml_key_always       = "always"
)

var (
	Never = DepthFunction{
		Enabled:  true,
		Function: gl.NEVER,
	}
	Less = DepthFunction{
		Enabled:  true,
		Function: gl.LESS,
	}
	Equal = DepthFunction{
		Enabled:  true,
		Function: gl.EQUAL,
	}
	LessEqual = DepthFunction{
		Enabled:  true,
		Function: gl.LEQUAL,
	}
	Greater = DepthFunction{
		Enabled:  true,
		Function: gl.GREATER,
	}
	NotEqual = DepthFunction{
		Enabled:  true,
		Function: gl.NOTEQUAL,
	}
	GreaterEqual = DepthFunction{
		Enabled:  true,
		Function: gl.GEQUAL,
	}
	Always = DepthFunction{
		Enabled:  true,
		Function: gl.ALWAYS,
	}
	DisabledDepth = DepthFunction{
		Enabled: false,
	}
)

type DepthFunction struct {
	Enabled  bool
	Function uint32
}

func (f *DepthFunction) Set() {
	if f.Enabled {
		gl.Enable(gl.DEPTH_TEST)
		gl.DepthFunc(f.Function)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
}

func (f *DepthFunction) UnmarshalYAML(value *yaml.Node) error {
	var tmp string
	value.Decode(&tmp)

	switch tmp {
	case yaml_key_never:
		*f = Never
	case yaml_key_less, Constants.Yaml_key_true:
		*f = Less
	case yaml_key_equal:
		*f = Equal
	case yaml_key_lessEqual:
		*f = LessEqual
	case yaml_key_greater:
		*f = Greater
	case yaml_key_notEqual:
		*f = NotEqual
	case yaml_key_greaterEqual:
		*f = GreaterEqual
	case yaml_key_always:
		*f = Always
	case Constants.Yaml_key_false:
		*f = DisabledDepth
	default:
		return fmt.Errorf("%s is not a depth function", tmp)
	}

	return nil
}

func GetCurrentDepthFunction() *DepthFunction {
	var enabled bool
	var function int32

	gl.GetBooleanv(gl.DEPTH_TEST, &enabled)
	gl.GetIntegerv(gl.DEPTH_FUNC, &function)

	return &DepthFunction{
		Enabled:  enabled,
		Function: uint32(function),
	}
}
