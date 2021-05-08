package main

import (
	"github.com/mattak/siga/cmd"
	"log"
	"math"
	"testing"
)

func TestMatrixGreaterEqual(t *testing.T) {
	matrix := cmd.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := matrix.GreaterEqual()
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
	if result[3] != 1 {
		log.Fatal("result[3] is not expected")
	}
	if result[4] != 0 {
		log.Fatal("result[4] is not expected")
	}
}
