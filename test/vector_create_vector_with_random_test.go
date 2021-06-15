package main

import (
	"github.com/mattak/siga/pkg"
	"testing"
)

func TestCreateVectorWithRandom(t *testing.T) {
	vector1 := pkg.CreateVectorWithRandom(3, 0, 1)
	vector2 := pkg.CreateVectorWithRandom(3, 0, 1)

	if len(vector1) != 3 || len(vector2) != 3 {
		t.Fatal("wrong size of vector")
	}
	ExpectRangeValue(t, "vector1[0]", vector1[0], 0.0, 1.0)
	ExpectRangeValue(t, "vector1[1]", vector1[1], 0.0, 1.0)
	ExpectRangeValue(t, "vector1[2]", vector1[2], 0.0, 1.0)
	ExpectRangeValue(t, "vector2[0]", vector2[0], 0.0, 1.0)
	ExpectRangeValue(t, "vector2[1]", vector2[1], 0.0, 1.0)
	ExpectRangeValue(t, "vector2[2]", vector2[2], 0.0, 1.0)
}
