package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestProfitFactor(t *testing.T) {
	if !math.IsInf((dataframe.Vector{1, 2, 3}).ProfitFactor(), 1) {
		t.Fatal("result should be infinite")
	}
	if (dataframe.Vector{-1, -2, -3}).ProfitFactor() != 0 {
		t.Fatal("result should be 0")
	}
	if (dataframe.Vector{-1, -2, 3}).ProfitFactor() != 1 {
		t.Fatal("result should be 0")
	}
}
