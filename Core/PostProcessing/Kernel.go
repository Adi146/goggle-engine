package PostProcessing

type Kernel struct {
	Kernel []float32 `yaml:"kernel"`

	Multiplier struct {
		Numerator   float32 `yaml:"numerator"`
		Denominator float32 `yaml:"denominator"`
	} `yaml:"kernelMultiplier"`

	Offset struct {
		Numerator   float32 `yaml:"numerator"`
		Denominator float32 `yaml:"denominator"`
	} `yaml:"kernelOffset"`
}

func (kernel *Kernel) GetOffset() float32 {
	return kernel.Offset.Numerator / kernel.Offset.Denominator
}

func (kernel *Kernel) GetKernel() []float32 {
	k := make([]float32, len(kernel.Kernel))
	mul := kernel.Multiplier.Numerator / kernel.Multiplier.Denominator

	for i := 0; i < len(kernel.Kernel); i++ {
		k[i] = mul * kernel.Kernel[i]
	}

	return k
}
