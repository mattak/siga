package main

import (
	"github.com/mattak/siga/pkg"
	"math"
	"testing"
)

func TestMatrixFilterIndexByOrNot(t *testing.T) {
	matrix := pkg.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 2, 1, 1},
	}
	result := matrix.FilterIndexByNotOr(1)
	ExpectInt(t, "result length is not expected", len(result), 2)
	ExpectInt(t, "result[0] is not expected", result[0], 0)
	ExpectInt(t, "result[2] is not expected", result[1], 2)
}
