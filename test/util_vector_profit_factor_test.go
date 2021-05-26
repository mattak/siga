package main

import (
	"github.com/mattak/siga/cmd"
	"math"
	"testing"
)

func TestProfitFactor(t *testing.T) {
	if !math.IsInf((cmd.Vector{1, 2, 3}).ProfitFactor(), 1) {
		t.Fatal("result should be infinite")
	}
	if (cmd.Vector{-1, -2, -3}).ProfitFactor() != 0 {
		t.Fatal("result should be 0")
	}
	if (cmd.Vector{-1, -2, 3}).ProfitFactor() != 1 {
		t.Fatal("result should be 0")
	}
}
