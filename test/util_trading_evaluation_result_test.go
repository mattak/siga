package main

import (
	"fmt"
	"github.com/mattak/siga/cmd"
	"log"
	"math"
	"testing"
)

func expectValue(t *testing.T, name string, real float64, expect float64) {
	if math.IsNaN(expect) && math.IsNaN(real) || math.Abs(real - expect) <= 1e-9 {
		return
	}
	t.Fatalf("%s is not matched: %f <=> %f", name, real, expect)
}

func expectAverageLoss(t *testing.T, name string, real cmd.AverageVariance, expect cmd.AverageVariance) {
	expectValue(t, fmt.Sprintf("%s.Total", name), expect.Total, real.Total)
	if expect.Count != real.Count {
		t.Fatalf("%s.Count is not matched: %d <=> %d", name, expect.Count, real.Count)
	}
	expectValue(t, fmt.Sprintf("%s.Average", name), expect.Average, real.Average)
	expectValue(t, fmt.Sprintf("%s.Variance", name), expect.Variance, real.Variance)
}

func TestTradingEvaluationResultCalculate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buy := cmd.Vector{}
		sell := cmd.Vector{}
		result, err := cmd.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		expectAverageLoss(t, "Gain", result.Gain, cmd.AverageVariance{0, 0, 0, 0})
		expectAverageLoss(t, "Loss", result.Loss, cmd.AverageVariance{0, 0, 0, 0})
		expectAverageLoss(t, "Profit", result.Profit, cmd.AverageVariance{0, 0, 0, 0})
		expectAverageLoss(t, "GainRate", result.GainRate, cmd.AverageVariance{0, 0, 0, 0})
		expectAverageLoss(t, "LossRate", result.LossRate, cmd.AverageVariance{0, 0, 0, 0})
		expectAverageLoss(t, "ProfitRate", result.ProfitRate, cmd.AverageVariance{0, 0, 0, 0})
		expectValue(t, "TotalBuy", result.TotalBuy, 0)
		expectValue(t, "TotalSell", result.TotalSell, 0)
		expectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		expectValue(t, "WinRate", result.WinRate, math.NaN())
		expectValue(t, "PayoffRatio", result.PayoffRatio, math.NaN())
		expectValue(t, "ProfitFactor", result.ProfitFactor, math.NaN())
		expectValue(t, "SharpRatio", result.SharpRatio, math.NaN())
		expectValue(t, "SortinoRatio", result.SortinoRatio, math.NaN())
		expectValue(t, "MaxDrawDown", result.MaxDrawDown, 0)
	})
	t.Run("length mismatch", func(t *testing.T) {
		buy := cmd.Vector{1}
		sell := cmd.Vector{}
		_, err := cmd.CreateTradingEvaluationResult(buy, sell)
		if err == nil {
			log.Fatal("error should be raised")
		}
	})
	t.Run("case1", func(t *testing.T) {
		buy := cmd.Vector{10, 10, 10}
		sell := cmd.Vector{10, 5, 15}

		result, err := cmd.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		expectAverageLoss(t, "Gain", result.Gain, cmd.AverageVariance{5, 2, 2.5, 2.5})
		expectAverageLoss(t, "Loss", result.Loss, cmd.AverageVariance{-5, 1, -5, 0})
		expectAverageLoss(t, "Profit", result.Profit, cmd.AverageVariance{0, 3, 0, math.Sqrt(50.0 / 3.0)})
		expectAverageLoss(t, "GainRate", result.GainRate, cmd.AverageVariance{0.5, 2, 0.25, 0.25})
		expectAverageLoss(t, "LossRate", result.LossRate, cmd.AverageVariance{-0.5, 1, -0.5, 0})
		expectAverageLoss(t, "ProfitRate", result.ProfitRate, cmd.AverageVariance{0, 3, 0, math.Sqrt(0.5 / 3.0)})
		expectValue(t, "TotalBuy", result.TotalBuy, 30)
		expectValue(t, "TotalSell", result.TotalSell, 30)
		expectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		expectValue(t, "WinRate", result.WinRate, 2.0/3.0)
		expectValue(t, "PayoffRatio", result.PayoffRatio, 0.5)
		expectValue(t, "ProfitFactor", result.ProfitFactor, 1)
		expectValue(t, "SharpRatio", result.SharpRatio, 0)
		expectValue(t, "SortinoRatio", result.SortinoRatio, math.NaN())
		expectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.5)
	})

	t.Run("case2", func(t *testing.T) {
		buy := cmd.Vector{10, 10, 10, 10}
		sell := cmd.Vector{8, 6, 12, 14}

		result, err := cmd.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		// 0.2^2 0.4^2 0.2^2 0.4^2
		// 0.04 0.16 0.04 0.16 => 0.32 + 0.08 = 0.4
		// sqrt(0.1)
		expectAverageLoss(t, "Gain", result.Gain, cmd.AverageVariance{6, 2, 3, 1})
		expectAverageLoss(t, "Loss", result.Loss, cmd.AverageVariance{-6, 2, -3, 1})
		expectAverageLoss(t, "Profit", result.Profit, cmd.AverageVariance{0, 4, 0, math.Sqrt(10.0)})
		expectAverageLoss(t, "GainRate", result.GainRate, cmd.AverageVariance{0.2 + 0.4, 2, 0.3, 0.1})
		expectAverageLoss(t, "LossRate", result.LossRate, cmd.AverageVariance{-0.2 - 0.4, 2, -0.3, 0.1})
		expectAverageLoss(t, "ProfitRate", result.ProfitRate, cmd.AverageVariance{0, 4, 0, math.Sqrt(0.1)})
		expectValue(t, "TotalBuy", result.TotalBuy, 40)
		expectValue(t, "TotalSell", result.TotalSell, 40)
		expectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		expectValue(t, "WinRate", result.WinRate, 0.5)
		expectValue(t, "PayoffRatio", result.PayoffRatio, 1)
		expectValue(t, "ProfitFactor", result.ProfitFactor, 1)
		expectValue(t, "SharpRatio", result.SharpRatio, 0)
		expectValue(t, "SortinoRatio", result.SortinoRatio, 0)
		expectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.4)
	})

	t.Run("case3", func(t *testing.T) {
		buy := cmd.Vector{10, 10, 10, 10, 20}
		sell := cmd.Vector{8, 6, 12, 14, 26}

		result, err := cmd.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		expectAverageLoss(t, "Gain", result.Gain, cmd.AverageVariance{12, 3, 4, math.Sqrt(8.0/3)})
		expectAverageLoss(t, "Loss", result.Loss, cmd.AverageVariance{-6, 2, -3, 1})
		expectAverageLoss(t, "Profit", result.Profit, cmd.AverageVariance{6, 5, 1.2, math.Sqrt(68.8/5)})
		expectAverageLoss(t, "GainRate", result.GainRate, cmd.AverageVariance{0.9, 3, 0.3, math.Sqrt(0.02/3.0)})
		expectAverageLoss(t, "LossRate", result.LossRate, cmd.AverageVariance{-0.2 - 0.4, 2, -0.3, 0.1})
		expectAverageLoss(t, "ProfitRate", result.ProfitRate, cmd.AverageVariance{0.3, 5, 0.06, math.Sqrt(0.472/5)})
		expectValue(t, "TotalBuy", result.TotalBuy, 60)
		expectValue(t, "TotalSell", result.TotalSell, 66)
		expectValue(t, "TotalReturnRate", result.TotalReturnRate, 0.1)
		expectValue(t, "WinRate", result.WinRate, 0.6)
		expectValue(t, "PayoffRatio", result.PayoffRatio, 4.0/3)
		expectValue(t, "ProfitFactor", result.ProfitFactor, 2)
		expectValue(t, "SharpRatio", result.SharpRatio, 0.06/math.Sqrt(0.472/5))
		expectValue(t, "SortinoRatio", result.SortinoRatio, 0.6)
		expectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.4)
	})
}
