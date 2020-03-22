package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"gopkg.in/yaml.v3"
)

const (
	FBO_type FrameBuffer.Type = "offscreen"
)

type OffScreenBuffer struct {
	FrameBuffer.FrameBuffer
	OffScreenTexture *Texture.Texture

	rbo uint32
}

func NewOffScreenBuffer(base FrameBuffer.FrameBuffer) (*OffScreenBuffer, error) {
	if err := gl.Init(); err != nil {
		return nil, err
	}

	offScreenTexture, err := NewOffscreenTexture(base.Viewport.Width, base.Viewport.Height)
	buff := OffScreenBuffer{
		FrameBuffer:      base,
		OffScreenTexture: offScreenTexture,
		rbo:              0,
	}
	if err != nil {
		return &buff, err
	}

	gl.GenRenderbuffers(1, &buff.rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buff.rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, buff.Viewport.Width, buff.Viewport.Height)
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)

	buff.Bind()
	defer buff.Unbind()
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, buff.OffScreenTexture.Target, buff.OffScreenTexture.ID, 0)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, buff.rbo)

	return &buff, buff.CheckCompleteness()
}

func (buff *OffScreenBuffer) Destroy() {
	buff.FrameBuffer.Destroy()
	gl.DeleteRenderbuffers(1, &buff.rbo)
	//buff.OffScreenTexture.Destroy()
}

func (buff *OffScreenBuffer) GetTextures() []Texture.ITexture {
	return []Texture.ITexture{
		buff.OffScreenTexture,
	}
}

func (buff *OffScreenBuffer) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		FrameBufferBase FrameBuffer.FrameBuffer `yaml:",inline"`
		Shaders         []Shader.Ptr            `yaml:"shaders"`
	}{
		FrameBufferBase: FrameBuffer.FrameBuffer{
			Type: FBO_type,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpBuffer, err := NewOffScreenBuffer(yamlConfig.FrameBufferBase)
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
