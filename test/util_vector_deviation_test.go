package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestVectorDeviation(t *testing.T) {
	data := cmd.Vector{1,2,3}

	// mean = 2, diff = [-1, 0, 1], deviation = 2 / 3
	d1 := data.Deviation(0, 3)
	if d1 >= 0.66 && d1 <= 0.67 {
		t.Fatal("deviation value is wrong")
	}
}
