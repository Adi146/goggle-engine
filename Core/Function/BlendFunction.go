package Function

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Utils/Constants"
	"github.com/go-gl/gl/v4.1-core/gl"
	"gopkg.in/yaml.v3"
)

var (
	DefaultBlend = BlendFunction{
		Enabled:     true,
		Source:      gl.SRC_ALPHA,
		Destination: gl.ONE_MINUS_SRC_ALPHA,
	}
	DisabledBlend = BlendFunction{
		Enabled: false,
	}
)

type BlendFunction struct {
	Enabled     bool
	Source      uint32
	Destination uint32
}

func (f *BlendFunction) Set() {
	if f.Enabled {
		gl.Enable(gl.BLEND)
		gl.BlendFunc(f.Source, f.Destination)
	} else {
		gl.Disable(gl.BLEND)
	}
}

func (f *BlendFunction) UnmarshalYAML(value *yaml.Node) error {
	var tmp string
	value.Decode(&tmp)

	switch tmp {
	case Constants.Yaml_key_true:
		*f = DefaultBlend
	case Constants.Yaml_key_false:
		*f = DisabledBlend
	default:
		return fmt.Errorf("%s is not a blend function", tmp)
	}

	return nil
}

func GetCurrentBlendFunction() *BlendFunction {
	var enabled bool
	var source int32
	var destination int32

	gl.GetBooleanv(gl.BLEND, &enabled)
	gl.GetIntegerv(gl.BLEND_SRC_ALPHA, &source)
	gl.GetIntegerv(gl.BLEND_DST_ALPHA, &destination)

	return &BlendFunction{
		Enabled:     enabled,
		Source:      uint32(source),
		Destination: uint32(destination),
	}
}
