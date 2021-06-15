package main

import (
	"github.com/mattak/siga/pkg"
	"math"
	"testing"
)

func TestHarmonicMean(t *testing.T) {
	t.Run("1,2,2", func(t *testing.T) {
		// 3/(1/1 + 1/2 + 1/2) = 3/2
		data := pkg.Vector{1, 2, 2}
		result := data.HarmonicMean(0, 3)
		ExpectValue(t, "result", result, 1.5)
	})
	t.Run("empty", func(t *testing.T) {
		data := pkg.Vector{}
		result := data.HarmonicMean(0, 0)
		ExpectValue(t, "result", result, math.NaN())
	})
}
