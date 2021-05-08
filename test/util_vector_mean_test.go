package main

import (
	"github.com/mattak/siga/cmd"
	"math"
	"testing"
)

func TestVectorMean(t *testing.T) {
	data := cmd.Vector{1, 2, 3}
	if data.Mean(0, 3) != 2 {
		t.Fatal("mean value is wrong")
	}
	if data.Mean(1, 2) != 2.5 {
		t.Fatal("mean value is wrong")
	}
	if data.Mean(2, 1) != 3 {
		t.Fatal("mean value is wrong")
	}
}

func TestVectorMeans(t *testing.T) {
	data := cmd.Vector{1, 2, 3}
	means := data.Means(2)
	if len(means) != 3 {
		t.Fatal("result length is wrong")
	}
	if means[0] != 1.5 {
		t.Fatal("mean[0] is wrong")
	}
	if means[1] != 2.5 {
		t.Fatal("mean[1] is wrong")
	}
	if !math.IsNaN(means[2]) {
		t.Fatal("mean[2] is wrong")
	}
}