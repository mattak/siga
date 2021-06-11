package main

import (
	"github.com/mattak/siga/toolkit"
	"log"
	"math"
	"testing"
)

func TestMatrixFilterNotIndex(t *testing.T) {
	matrix := toolkit.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := matrix.FilterNotIndex()
	if len(result) != 3 {
		log.Fatal("result length is not expected")
	}
	if result[0] != 1 {
		log.Fatal("result[0] is not expected")
	}
	if result[1] != 2 {
		log.Fatal("result[1] is not expected")
	}
	if result[2] != 4 {
		log.Fatal("result[2] is not expected")
	}
}
