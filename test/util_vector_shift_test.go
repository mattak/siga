package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestVectorShift(t *testing.T) {
	data := dataframe.Vector{1,2,3}
	r1 := data.Shift(0)
	if r1[0] != data[0] || r1[1] != data[1] || r1[2] != data[2] {
		t.Fatal("r1 shift value is wrong")
	}
	r2 := data.Shift(1)
	if !math.IsNaN(r2[0]) || r2[1] != data[0] || r2[2] != data[1] {
		t.Fatal("r2 shift value is wrong: ", math.IsNaN(r2[0]))
	}
	r3 := data.Shift(-1)
	if r3[0] != data[1] || r3[1] != data[2] || !math.IsNaN(r3[2]) {
		t.Fatal("r3 shift value is wrong")
	}
}
