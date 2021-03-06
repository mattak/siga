package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestDollarCostAverage(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		data := dataframe.Vector{1, 2, 3}

		v := data.DollarCostAverage()
		ExpectInt(t, "len(v)", len(v), 3)
		ExpectValue(t, "v[0]", v[0], 1)
		ExpectValue(t, "v[1]", v[1], 4.0/3)
		ExpectValue(t, "v[2]", v[2], 3.0/(1.5+1/3.0))
	})

	t.Run("nan1", func(t *testing.T) {
		data := dataframe.Vector{1, math.NaN(), 3}

		v := data.DollarCostAverage()
		ExpectInt(t, "len(v)", len(v), 3)
		ExpectValue(t, "v[0]", v[0], 1)
		ExpectValue(t, "v[1]", v[1], 1)
		ExpectValue(t, "v[2]", v[2], 2.0/(1+1.0/3))
	})

	t.Run("zero1", func(t *testing.T) {
		data := dataframe.Vector{1, 0, 3}

		v := data.DollarCostAverage()
		ExpectInt(t, "len(v)", len(v), 3)
		ExpectValue(t, "v[0]", v[0], 1)
		ExpectValue(t, "v[1]", v[1], 1)
		ExpectValue(t, "v[2]", v[2], 2.0/(1+1.0/3))
	})

	t.Run("nan0", func(t *testing.T) {
		data := dataframe.Vector{math.NaN(), 2, 3}

		v := data.DollarCostAverage()
		ExpectInt(t, "len(v)", len(v), 3)
		ExpectValue(t, "v[0]", v[0], math.NaN())
		ExpectValue(t, "v[1]", v[1], 2)
		ExpectValue(t, "v[2]", v[2], 2.0/(0.5+1.0/3))
	})

	t.Run("zero0", func(t *testing.T) {
		data := dataframe.Vector{0, 2, 3}

		v := data.DollarCostAverage()
		ExpectInt(t, "len(v)", len(v), 3)
		ExpectValue(t, "v[0]", v[0], math.NaN())
		ExpectValue(t, "v[1]", v[1], 2)
		ExpectValue(t, "v[2]", v[2], 2.0/(0.5+1.0/3))
	})
}
