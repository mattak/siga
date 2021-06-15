package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestMartinegaleMean(t *testing.T) {
	t.Run("1,2,3", func(t *testing.T) {
		data := dataframe.Vector{1, 2, 3}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 1)
	})
	t.Run("3,2,1", func(t *testing.T) {
		// last:    3, 3, 2,
		// step:    1, 2, 4,
		// invests: 1, 3, 7,
		// amount:  1/3, 1/3+2/2, 1/3+2/2+4/1 = 5+1/3 = 16/3
		// result = invests/amount = 7/(16/3) = 21/16
		data := dataframe.Vector{3, 2, 1}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 21.0/16.0)
	})
	t.Run("3,3,3", func(t *testing.T) {
		data := dataframe.Vector{3, 3, 3}
		result := data.MartinegaleMean(0, 3, 1)
		ExpectValue(t, "result", result, 3)
	})
	t.Run("empty", func(t *testing.T) {
		data := dataframe.Vector{}
		result := data.MartinegaleMean(0, 0, 1)
		ExpectValue(t, "result", result, math.NaN())
	})
}
