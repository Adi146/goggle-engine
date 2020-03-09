package GeometryMath_test

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"testing"
)

func TestCross(t *testing.T) {
	a := &GeometryMath.Vector3{1, 2, 3}
	b := &GeometryMath.Vector3{-7, 8, 9}
	result := a.Cross(b)
	expectedResult := &GeometryMath.Vector3{-6, -30, 22}

	if *result != *expectedResult {
		t.Errorf("Cross(%f, %f) not matching (expected %f, got %f)", *a, *b, *expectedResult, *result)
	}
}
