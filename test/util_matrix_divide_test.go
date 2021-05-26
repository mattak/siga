package main

import (
	"github.com/mattak/siga/cmd"
	"log"
	"math"
	"testing"
)

func TestMatrixDivide(t *testing.T) {
	matrix := cmd.Matrix{
		{math.NaN(), 1, math.NaN(), 1, 0, 1, 1, 1},
		{math.NaN(), math.NaN(), 1, 0, 1, 1, 2, 0.5},
	}
	result := matrix.Divide()
	if len(result) != 8 {
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
	if !math.IsNaN(result[3]) {
		log.Fatal("result[3] is not expected")
	}
	if result[4] != 0 {
		log.Fatal("result[4] is not expected")
	}
	if result[5] != 1.0 {
		log.Fatal("result[5] is not expected")
	}
	if result[6] != 0.5 {
		log.Fatal("result[6] is not expected")
	}
	if result[7] != 2 {
		log.Fatal("result[7] is not expected")
	}
}
