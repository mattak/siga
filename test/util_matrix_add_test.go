package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"math"
	"testing"
)

func TestMatrixAdd(t *testing.T) {
	matrix := dataframe.Matrix{
		{math.NaN(), math.NaN(), 1, 1, 2},
		{math.NaN(), 1, math.NaN(), 0, 1},
	}
	result := matrix.Add()
	if len(result) != 5 {
		log.Fatal("result length is not expected")
	}
	if !math.IsNaN(result[0]) {
		log.Fatal("result[0] is not expected")
	}
	if !math.IsNaN(result[1]) {
		log.Fatal("result[1] is not expected")
	}
	if !math.IsNaN(result[2]) {
		log.Fatal("result[2] is not expected")
	}
	if result[3] != 1 {
		log.Fatal("result[3] is not expected")
	}
	if result[4] != 3 {
		log.Fatal("result[4] is not expected")
	}
}
