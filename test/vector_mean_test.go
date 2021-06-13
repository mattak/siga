package main

import (
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	t.Run("1,2,3", func(t *testing.T) {
		data := toolkit.Vector{1, 2, 3}
		result := data.Mean(0, 3)
		ExpectValue(t, "result", result, 2)
	})
	t.Run("empty", func(t *testing.T) {
		data := toolkit.Vector{}
		result := data.Mean(0, 0)
		ExpectValue(t, "result", result, math.NaN())
	})
}
