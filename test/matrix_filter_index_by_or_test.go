package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestMatrixFilterIndexByOr(t *testing.T) {
	matrix := dataframe.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 2, 1, 1},
	}
	result := matrix.FilterIndexByOr(1)
	ExpectInt(t, "result length is not expected", len(result), 3)
	ExpectInt(t, "result[0] is not expected", result[0], 1)
	ExpectInt(t, "result[1] is not expected", result[1], 3)
	ExpectInt(t, "result[2] is not expected", result[2], 4)
}
