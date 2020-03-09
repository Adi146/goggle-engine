package GeometryMath_test

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"testing"
)

func TestMul(t *testing.T) {
	a := &GeometryMath.Matrix4x4{
		{5, 7, 9, 10},
		{2, 3, 3, 8},
		{8, 10, 2, 3},
		{3, 3, 4, 8},
	}

	b := &GeometryMath.Matrix4x4{
		{3, 10, 12, 18},
		{12, 1, 4, 9},
		{9, 10, 12, 2},
		{3, 12, 4, 10},
	}

	result := a.Mul(b)

	expectedResult := &GeometryMath.Matrix4x4{
		{210, 267, 236, 271},
		{93, 149, 104, 149},
		{171, 146, 172, 268},
		{105, 169, 128, 169},
	}

	if !result.Equals(expectedResult, 1e-5) {
		t.Errorf("Mul(%f, %f) not matching (expecting %f, got %f)", *a, *b, *expectedResult, *result)
	}
}

func TestInverse(t *testing.T) {
	a := &GeometryMath.Matrix4x4{
		{-3, 0, 3, -1},
		{0, 3, 4, -1},
		{-4, -2, 2, -4},
		{2, 0, 1, 11},
	}

	result := a.Inverse()

	expectedResult := &GeometryMath.Matrix4x4{
		{-21.0 / 19.0, 17.0 / 38.0, 51.0 / 76.0, 7.0 / 38.0},
		{1.0, -1.0 / 4.0, -7.0 / 8.0, -1.0 / 4.0},
		{-13.0 / 19.0, 31.0 / 76.0, 93.0 / 152.0, 15.0 / 76.0},
		{5.0 / 19.0, -9.0 / 76.0, -27.0 / 152.0, 3.0 / 76.0},
	}

	if !result.Equals(expectedResult, 1e-5) {
		t.Errorf("Inv(\n\t%f\n) not mathing (expected \n\t%f\n, got \n\t%f\n)", *a, *expectedResult, *result)
	}
}
