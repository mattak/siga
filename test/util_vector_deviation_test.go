package main

import (
	"github.com/mattak/siga/cmd"
	"log"
	"math"
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

func TestVectorDeviations(t *testing.T) {
	data := cmd.Vector{1,2,3}

	// mean = 2, diff = [-1, 0, 1], deviation = 2 / 3
	deviations := data.Deviations(2)
	if len(deviations) != 3 {
		log.Fatal("deviations length is not expected")
	}
	// mean: 1.5, diff = [-0.5, 0.5], diff_2 = [0.25, 0.25], sum_diff2 = 0.5, d2 = 0.25
	if deviations[0] != 0.5 {
		log.Fatal("deviations[0] is not expected")
	}
	if deviations[1] != 0.5 {
		log.Fatal("deviations[1] is not expected")
	}
	if !math.IsNaN(deviations[2]) {
		log.Fatal("deviations[2] is not expected")
	}
}
