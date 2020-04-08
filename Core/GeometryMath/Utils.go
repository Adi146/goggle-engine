package GeometryMath

import "math"

func Abs(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

func Sin(a float32) float32 {
	return float32(math.Sin(float64(a)))
}

func Cos(a float32) float32 {
	return float32(math.Cos(float64(a)))
}

func Tan(a float32) float32 {
	return float32(math.Tan(float64(a)))
}

func ASin(a float32) float32 {
	return float32(math.Asin(float64(a)))
}

func ACos(a float32) float32 {
	return float32(math.Acos(float64(a)))
}

func ATan2(y float32, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func Pow(a float32, exp float32) float32 {
	return float32(math.Pow(float64(a), float64(exp)))
}

func Sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func Equals(f1 float32, f2 float32, threshold float32) bool {
	return Abs(f1-f2) < threshold
}

func Min(x float32, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

func Max(x float32, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}