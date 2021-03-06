package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestMatrixFilterIndexByAnd(t *testing.T) {
	m := dataframe.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := m.FilterIndexByAnd(1)
	ExpectInt(t, "result length is not expected", len(result), 1)
	ExpectInt(t, "result[0] is not expected", result[0], 3)
}
