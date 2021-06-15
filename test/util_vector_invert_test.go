package main

import (
	"github.com/mattak/siga/pkg"
	"math"
	"testing"
)

func TestVectorInvert(t *testing.T) {
	data := pkg.Vector{1,0, math.NaN()}
	r := data.Invert()
	if r[0] != 0 || r[1] != 1 || !math.IsNaN(r[2]) {
		t.Fatal("r invert value is wrong")
	}
}
