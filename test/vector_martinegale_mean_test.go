package main

import (
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func TestMartinegaleMean(t *testing.T) {
	t.Run("1,2,3", func(t *testing.T) {
		data := toolkit.Vector{1, 2, 3}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 1)
	})
	t.Run("3,2,1", func(t *testing.T) {
		// cost: 1, 2, 4
		// amount: 1, 1, 4
		// 7/6
		data := toolkit.Vector{3, 2, 1}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 7.0/6.0)
	})
	t.Run("3,3,3", func(t *testing.T) {
		data := toolkit.Vector{3, 3, 3}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 3)
	})
	t.Run("empty", func(t *testing.T) {
		data := toolkit.Vector{}
		result := data.MartinegaleMean(0, 0, 1)
		ExpectValue(t, "result", result, math.NaN())
	})
}
