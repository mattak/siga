package main

import (
	"github.com/mattak/siga/cmd"
	"math"
	"testing"
)

func TestVectorSigmaAnomalies(t *testing.T) {
	t.Run("1 seq", func(t *testing.T) {
		data := cmd.Vector{1, 2, 3}
		data.Reverse()
		result := data.SigmaAnomalies(1, 1)
		result.Reverse()
		if len(result) != 3 {
			t.Fatal("result length is wrong")
		}
		if !math.IsNaN(result[0]) {
			t.Fatal("result[0] should be NaN")
		}
		// mean = 1, deviation = 0, diff = 1, deviation*thresholdSigma = 0,
		if result[1] != 0.0 {
			t.Fatalf("result[1] should be 0, but %f", result[1])
		}
		if result[2] != 0.0 {
			t.Fatalf("result[2] should be 0, but %f", result[2])
		}
	})

	t.Run("2 seq", func(t *testing.T){
		data := cmd.Vector{1, 2, 3}
		data.Reverse()
		result := data.SigmaAnomalies(2, 1)
		result.Reverse()
		if len(result) != 3 {
			t.Fatal("result length is wrong")
		}
		if !math.IsNaN(result[0]) {
			t.Fatal("result[0] should be NaN")
		}
		if !math.IsNaN(result[1]) {
			t.Fatal("result[1] should be NaN")
		}
		// mean = 1.5, deviation = 0.5, diff = 1.5, deviation*thresholdSigma = 0.5,
		if result[2] != 3.0 {
			t.Fatalf("result[2] should be 3.0, but %f", result[2])
		}
	})
}
