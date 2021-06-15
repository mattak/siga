package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"math"
	"testing"
)

func TestMatrixGreaterThan(t *testing.T) {
	matrix := dataframe.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := matrix.GreaterThan()
	if len(result) != 5 {
		log.Fatal("result length is not expected")
	}
	if !math.IsNaN(result[0]) {
		log.Fatal("result[0] is not expected")
	}
	if !math.IsNaN(result[1]) {
		log.Fatal("result[1] is not expected")
	}
	if result[2] != 1 {
		log.Fatal("result[2] is not expected")
	}
	if result[3] != 0 {
		log.Fatal("result[3] is not expected")
	}
	if result[4] != 0 {
		log.Fatal("result[4] is not expected")
	}
}
