package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestValueMean(t *testing.T) {
	t.Run("1,2,2", func(t *testing.T) {
		// price:1, amounts:0, invests:0, expect:1, curr:0 -> amount:1, invest: 1
		// price:2, amounts:1, invests:1, expect:2, curr:2 -> amount:0, invest: 0
		// price:2, amounts:1, invests:1, expect:3, curr:2 -> amount:0.5, invest: 1
		// result, amounts:1.5, invests:2, price: invests/amounts = 2/1.5 = 1.3333
		data := dataframe.Vector{1, 2, 2}
		result := data.ValueMean(0, 3)
		ExpectValue(t, "result", result, 4.0/3)
	})
	t.Run("empty", func(t *testing.T) {
		data := dataframe.Vector{}
		result := data.ValueMean(0, 0)
		ExpectValue(t, "result", result, math.NaN())
	})
}
