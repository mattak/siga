package main

import (
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func TestMatrixFilterNotIndex(t *testing.T) {
	matrix := toolkit.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := matrix.FilterIndexByNotAnd(1)
	ExpectInt(t, "result length is not expected", len(result), 4)
	ExpectInt(t, "result[0] is not expected", result[0], 0)
	ExpectInt(t, "result[1] is not expected", result[1], 1)
	ExpectInt(t, "result[2] is not expected", result[2], 2)
	ExpectInt(t, "result[3] is not expected", result[3], 4)
}
