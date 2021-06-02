package main

import (
	"github.com/mattak/siga/cmd"
	"math"
	"testing"
)

func TestPayoffRatio(t *testing.T) {
	result1 := (cmd.Vector{1, 2, 3}).PayoffRatio()
	if !math.IsNaN(result1.PayoffRatio) {
		t.Fatal("result1 payoff ratio is wrong", result1.PayoffRatio)
	}
	if result1.ProfitCount != 3 || result1.ProfitAverage != 2 {
		t.Fatal("result1 profits is wrong")
	}
	if result1.LossCount != 0 || !math.IsNaN(result1.LossAverage) {
		t.Fatal("result1 loss is wrong")
	}

	result2 := (cmd.Vector{-1, -2, -3}).PayoffRatio()
	if !math.IsNaN(result2.PayoffRatio) {
		t.Fatal("result2 payoff ratio is wrong")
	}
	if result2.ProfitCount != 0 || !math.IsNaN(result2.ProfitAverage) {
		t.Fatal("result2 profit is wrong")
	}
	if result2.LossCount != 3 || result2.LossAverage != -2 {
		t.Fatal("result2 loss is wrong", result2.LossAverage)
	}

	result3 := (cmd.Vector{2, -1}).PayoffRatio()
	if result3.PayoffRatio != 2 {
		t.Fatal("result3 payoff ratio is wrong")
	}
	if result3.ProfitCount != 1 || result3.ProfitAverage != 2 {
		t.Fatal("result3 profit is wrong")
	}
	if result3.LossCount != 1 || result3.LossAverage != -1 {
		t.Fatal("result3 loss is wrong")
	}

	result4 := (cmd.Vector{1, 1, -1}).PayoffRatio()
	if result4.PayoffRatio != 1 {
		t.Fatal("result4 payoff ratio is wrong")
	}
	if result4.ProfitCount != 2 || result4.ProfitAverage != 1 {
		t.Fatal("result4 profit is wrong")
	}
	if result4.LossCount != 1 || result4.LossAverage != -1 {
		t.Fatal("result4 loss is wrong")
	}
}
