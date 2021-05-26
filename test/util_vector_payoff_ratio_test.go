package main

import (
	"github.com/mattak/siga/cmd"
	"math"
	"testing"
)

func TestPayoffRatio(t *testing.T) {
	if !math.IsNaN((cmd.Vector{1, 2, 3}).PayoffRatio()) {
		t.Fatal("result should be nan")
	}
	if !math.IsNaN((cmd.Vector{-1, -2, -3}).PayoffRatio()) {
		t.Fatal("result should be nan")
	}
	if (cmd.Vector{2, -1}).PayoffRatio() != 2 {
		t.Fatal("result should be 0: ")
	}
	if (cmd.Vector{1, 2, -3}).ProfitFactor() != 1 {
		t.Fatal("result should be 0")
	}
}
