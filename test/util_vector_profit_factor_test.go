package main

import (
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func TestProfitFactor(t *testing.T) {
	if !math.IsInf((toolkit.Vector{1, 2, 3}).ProfitFactor(), 1) {
		t.Fatal("result should be infinite")
	}
	if (toolkit.Vector{-1, -2, -3}).ProfitFactor() != 0 {
		t.Fatal("result should be 0")
	}
	if (toolkit.Vector{-1, -2, 3}).ProfitFactor() != 1 {
		t.Fatal("result should be 0")
	}
}
