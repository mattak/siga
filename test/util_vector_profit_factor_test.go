package main

import (
	"github.com/mattak/siga/pkg"
	"math"
	"testing"
)

func TestProfitFactor(t *testing.T) {
	if !math.IsInf((pkg.Vector{1, 2, 3}).ProfitFactor(), 1) {
		t.Fatal("result should be infinite")
	}
	if (pkg.Vector{-1, -2, -3}).ProfitFactor() != 0 {
		t.Fatal("result should be 0")
	}
	if (pkg.Vector{-1, -2, 3}).ProfitFactor() != 1 {
		t.Fatal("result should be 0")
	}
}
