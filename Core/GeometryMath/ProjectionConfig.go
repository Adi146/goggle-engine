package GeometryMath

type IProjectionConfig interface {
	IMatrixConfig
	GetNear() float32
	GetFar() float32
	GetPlane(distance float32) (float32, float32)
	SetSize(width, height, near, far float32)
}

type OrthographicConfig struct {
	Left   float32 `yaml:"left"`
	Right  float32 `yaml:"right"`
	Bottom float32 `yaml:"bottom"`
	Top    float32 `yaml:"top"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (config *OrthographicConfig) GetNear() float32 {
	return config.Near
}

func (config *OrthographicConfig) GetFar() float32 {
	return config.Far
}

func (config *OrthographicConfig) Decode() Matrix4x4 {
	return Orthographic(
		config.Left,
		config.Right,
		config.Bottom,
		config.Top,
		config.Near,
		config.Far,
	)
}

func (config OrthographicConfig) SetSize(width, height, near, far float32) {
	halfWidth := width / 2.0
	halfHeight := height / 2.0

	config.Left = -halfWidth
	config.Right = halfWidth
	config.Bottom = -halfHeight
	config.Top = halfHeight
	config.Near = near
	config.Far = far
}

func (config *OrthographicConfig) GetPlane(distance float32) (float32, float32) {
	height := config.Top - config.Bottom
	width := config.Right - config.Left

	return width, height
}

type PerspectiveConfig struct {
	Fovy   float32 `yaml:"fovy"`
	Aspect float32 `yaml:"aspect"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (config *PerspectiveConfig) Decode() Matrix4x4 {
	return Perspective(
		Radians(config.Fovy),
		config.Aspect,
		config.Near,
		config.Far,
	)
}

func (config *PerspectiveConfig) GetNear() float32 {
	return config.Near
}

func (config *PerspectiveConfig) GetFar() float32 {
	return config.Far
}

func (config *PerspectiveConfig) GetPlane(distance float32) (float32, float32) {
	tanY := Tan(Radians(config.Fovy * 0.5))

	height := distance * tanY
	width := height * config.Aspect

	return width, height
}

func (config *PerspectiveConfig) SetSize(width, height, near, far float32) {
	config.Aspect = width / height
	config.Near = near
	config.Far = far
}
