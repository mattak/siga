package main

import (
	"github.com/mattak/siga/pkg"
	"log"
	"math"
	"testing"
)

func TestTradingEvaluationResultCalculate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buy := pkg.Vector{}
		sell := pkg.Vector{}
		result, err := pkg.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		ExpectAverageLoss(t, "Gain", result.Gain, pkg.AverageVariance{0, 0, 0, 0})
		ExpectAverageLoss(t, "Loss", result.Loss, pkg.AverageVariance{0, 0, 0, 0})
		ExpectAverageLoss(t, "Profit", result.Profit, pkg.AverageVariance{0, 0, 0, 0})
		ExpectAverageLoss(t, "GainRate", result.GainRate, pkg.AverageVariance{0, 0, 0, 0})
		ExpectAverageLoss(t, "LossRate", result.LossRate, pkg.AverageVariance{0, 0, 0, 0})
		ExpectAverageLoss(t, "ProfitRate", result.ProfitRate, pkg.AverageVariance{0, 0, 0, 0})
		ExpectValue(t, "TotalBuy", result.TotalBuy, 0)
		ExpectValue(t, "TotalSell", result.TotalSell, 0)
		ExpectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		ExpectValue(t, "WinRate", result.WinRate, math.NaN())
		ExpectValue(t, "PayoffRatio", result.PayoffRatio, math.NaN())
		ExpectValue(t, "ProfitFactor", result.ProfitFactor, math.NaN())
		ExpectValue(t, "SharpRatio", result.SharpRatio, math.NaN())
		ExpectValue(t, "SortinoRatio", result.SortinoRatio, math.NaN())
		ExpectValue(t, "MaxDrawDown", result.MaxDrawDown, 0)
	})
	t.Run("length mismatch", func(t *testing.T) {
		buy := pkg.Vector{1}
		sell := pkg.Vector{}
		_, err := pkg.CreateTradingEvaluationResult(buy, sell)
		if err == nil {
			log.Fatal("error should be raised")
		}
	})
	t.Run("case1", func(t *testing.T) {
		buy := pkg.Vector{10, 10, 10}
		sell := pkg.Vector{10, 5, 15}

		result, err := pkg.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		ExpectAverageLoss(t, "Gain", result.Gain, pkg.AverageVariance{5, 2, 2.5, 2.5})
		ExpectAverageLoss(t, "Loss", result.Loss, pkg.AverageVariance{-5, 1, -5, 0})
		ExpectAverageLoss(t, "Profit", result.Profit, pkg.AverageVariance{0, 3, 0, math.Sqrt(50.0 / 3.0)})
		ExpectAverageLoss(t, "GainRate", result.GainRate, pkg.AverageVariance{0.5, 2, 0.25, 0.25})
		ExpectAverageLoss(t, "LossRate", result.LossRate, pkg.AverageVariance{-0.5, 1, -0.5, 0})
		ExpectAverageLoss(t, "ProfitRate", result.ProfitRate, pkg.AverageVariance{0, 3, 0, math.Sqrt(0.5 / 3.0)})
		ExpectValue(t, "TotalBuy", result.TotalBuy, 30)
		ExpectValue(t, "TotalSell", result.TotalSell, 30)
		ExpectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		ExpectValue(t, "WinRate", result.WinRate, 2.0/3.0)
		ExpectValue(t, "PayoffRatio", result.PayoffRatio, 0.5)
		ExpectValue(t, "ProfitFactor", result.ProfitFactor, 1)
		ExpectValue(t, "SharpRatio", result.SharpRatio, 0)
		ExpectValue(t, "SortinoRatio", result.SortinoRatio, math.NaN())
		ExpectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.5)
	})

	t.Run("case2", func(t *testing.T) {
		buy := pkg.Vector{10, 10, 10, 10}
		sell := pkg.Vector{8, 6, 12, 14}

		result, err := pkg.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		// 0.2^2 0.4^2 0.2^2 0.4^2
		// 0.04 0.16 0.04 0.16 => 0.32 + 0.08 = 0.4
		// sqrt(0.1)
		ExpectAverageLoss(t, "Gain", result.Gain, pkg.AverageVariance{6, 2, 3, 1})
		ExpectAverageLoss(t, "Loss", result.Loss, pkg.AverageVariance{-6, 2, -3, 1})
		ExpectAverageLoss(t, "Profit", result.Profit, pkg.AverageVariance{0, 4, 0, math.Sqrt(10.0)})
		ExpectAverageLoss(t, "GainRate", result.GainRate, pkg.AverageVariance{0.2 + 0.4, 2, 0.3, 0.1})
		ExpectAverageLoss(t, "LossRate", result.LossRate, pkg.AverageVariance{-0.2 - 0.4, 2, -0.3, 0.1})
		ExpectAverageLoss(t, "ProfitRate", result.ProfitRate, pkg.AverageVariance{0, 4, 0, math.Sqrt(0.1)})
		ExpectValue(t, "TotalBuy", result.TotalBuy, 40)
		ExpectValue(t, "TotalSell", result.TotalSell, 40)
		ExpectValue(t, "TotalReturnRate", result.TotalReturnRate, 0)
		ExpectValue(t, "WinRate", result.WinRate, 0.5)
		ExpectValue(t, "PayoffRatio", result.PayoffRatio, 1)
		ExpectValue(t, "ProfitFactor", result.ProfitFactor, 1)
		ExpectValue(t, "SharpRatio", result.SharpRatio, 0)
		ExpectValue(t, "SortinoRatio", result.SortinoRatio, 0)
		ExpectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.4)
	})

	t.Run("case3", func(t *testing.T) {
		buy := pkg.Vector{10, 10, 10, 10, 20}
		sell := pkg.Vector{8, 6, 12, 14, 26}

		result, err := pkg.CreateTradingEvaluationResult(buy, sell)
		if err != nil {
			t.Fatal("error should not be raised")
		}

		ExpectAverageLoss(t, "Gain", result.Gain, pkg.AverageVariance{12, 3, 4, math.Sqrt(8.0/3)})
		ExpectAverageLoss(t, "Loss", result.Loss, pkg.AverageVariance{-6, 2, -3, 1})
		ExpectAverageLoss(t, "Profit", result.Profit, pkg.AverageVariance{6, 5, 1.2, math.Sqrt(68.8/5)})
		ExpectAverageLoss(t, "GainRate", result.GainRate, pkg.AverageVariance{0.9, 3, 0.3, math.Sqrt(0.02/3.0)})
		ExpectAverageLoss(t, "LossRate", result.LossRate, pkg.AverageVariance{-0.2 - 0.4, 2, -0.3, 0.1})
		ExpectAverageLoss(t, "ProfitRate", result.ProfitRate, pkg.AverageVariance{0.3, 5, 0.06, math.Sqrt(0.472/5)})
		ExpectValue(t, "TotalBuy", result.TotalBuy, 60)
		ExpectValue(t, "TotalSell", result.TotalSell, 66)
		ExpectValue(t, "TotalReturnRate", result.TotalReturnRate, 0.1)
		ExpectValue(t, "WinRate", result.WinRate, 0.6)
		ExpectValue(t, "PayoffRatio", result.PayoffRatio, 4.0/3)
		ExpectValue(t, "ProfitFactor", result.ProfitFactor, 2)
		ExpectValue(t, "SharpRatio", result.SharpRatio, 0.06/math.Sqrt(0.472/5))
		ExpectValue(t, "SortinoRatio", result.SortinoRatio, 0.6)
		ExpectValue(t, "MaxDrawDown", result.MaxDrawDown, -0.4)
	})
}
