package ShadowMapping

import (
	core "github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"gopkg.in/yaml.v3"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	FBO_type = "shadowMap"
)

type FrameBuffer struct {
	core.FrameBuffer
	ShadowMap *Texture.Texture
}

func NewFrameBuffer(base core.FrameBuffer) (*FrameBuffer, error) {
	if err := gl.Init(); err != nil {
		return nil, err
	}

	shadowMap, err := NewShadowMap(base.Viewport.Width, base.Viewport.Height)
	buff := FrameBuffer{
		FrameBuffer: base,
		ShadowMap:   shadowMap,
	}
	if err != nil {
		return &buff, err
	}

	buff.Bind()
	defer buff.Unbind()
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, buff.ShadowMap.Target, buff.ShadowMap.ID, 0)
	gl.DrawBuffer(gl.NONE)
	gl.ReadBuffer(gl.NONE)

	return &buff, buff.CheckCompleteness()
}

func (buff *FrameBuffer) Destroy() {
	buff.FrameBuffer.Destroy()
}

func (buff *FrameBuffer) GetTextures() []Texture.ITexture {
	return []Texture.ITexture{
		buff.ShadowMap,
	}
}

func (buff *FrameBuffer) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		FrameBufferBase core.FrameBuffer `yaml:",inline"`
		Shaders         []Shader.Ptr     `yaml:"shaders"`
	}{
		FrameBufferBase: core.FrameBuffer{
			Type: FBO_type,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpBuffer, err := NewFrameBuffer(yamlConfig.FrameBufferBase)
	if err != nil {
		return err
	}

	for _, shader := range yamlConfig.Shaders {
		if err := shader.BindObject(tmpBuffer); err != nil {
			return err
		}
	}

	*buff = *tmpBuffer

	return nil
}
