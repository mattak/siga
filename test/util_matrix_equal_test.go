package main

import (
	"github.com/mattak/siga/pkg"
	"log"
	"math"
	"testing"
)

func TestMatrixEqual(t *testing.T) {
	matrix := pkg.Matrix{
		{math.NaN(), math.NaN(), 1, 0, 1},
		{math.NaN(), 1, 1, 1, 1},
		{math.NaN(), 1, 0, 1, 1},
	}
	result := matrix.Equal()
	if len(result) != 5 {
		log.Fatal("result length is not expected")
	}
	if result[0] != 1.0 {
		log.Fatal("result[0] is not expected")
	}
	if result[1] != 0.0 {
		log.Fatal("result[1] is not expected")
	}
	if result[2] != 0.0 {
		log.Fatal("result[2] is not expected")
	}
	if result[3] != 0.0 {
		log.Fatal("result[3] is not expected")
	}
	if result[4] != 1.0 {
		log.Fatal("result[4] is not expected")
	}
}
