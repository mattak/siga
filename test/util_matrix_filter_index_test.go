package main

import (
	"github.com/mattak/siga/cmd"
	"log"
	"math"
	"testing"
)

func TestMatrixFilterIndex(t *testing.T) {
	matrix := cmd.Matrix{
		{math.NaN(), math.NaN(), 2, 1, 0},
		{math.NaN(), 1, 1, 1, 1},
	}
	result := matrix.FilterIndex()
	if len(result) != 2 {
		log.Fatal("result length is not expected")
	}
	if result[0] != 0 {
		log.Fatal("result[0] is not expected")
	}
	if result[1] != 3 {
		log.Fatal("result[1] is not expected")
	}
}
