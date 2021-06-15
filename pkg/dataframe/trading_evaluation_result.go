package dataframe

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type TradingEvaluationResult struct {
	// Detail
	Gain       AverageVariance `json:"gain"`
	Loss       AverageVariance `json:"loss"`
	Profit     AverageVariance `json:"profit"`
	GainRate   AverageVariance `json:"gain_rate"`
	LossRate   AverageVariance `json:"loss_rate"`
	ProfitRate AverageVariance `json:"profit_rate"`

	// Total
	TotalBuy        float64 `json:"total_buy"`
	TotalSell       float64 `json:"total_sell"`
	TotalReturnRate float64 `json:"total_return_rate"`

	// Trade
	WinRate      float64 `json:"win_rate"`
	PayoffRatio  float64 `json:"payoff_ratio"`
	ProfitFactor float64 `json:"profit_factor"`

	// Risk
	SharpRatio   float64 `json:"sharp_ratio"`
	SortinoRatio float64 `json:"sortino_ratio"`
	MaxDrawDown  float64 `json:"max_draw_down"`
}

func (pr *TradingEvaluationResult) ToVector(isDetailOutput bool) Vector {
	if isDetailOutput {
		return Vector{
			// Detail
			pr.Gain.Total,
			float64(pr.Gain.Count),
			pr.Gain.Average,
			pr.Gain.Variance,

			pr.Loss.Total,
			float64(pr.Loss.Count),
			pr.Loss.Average,
			pr.Loss.Variance,

			pr.Profit.Total,
			float64(pr.Profit.Count),
			pr.Profit.Average,
			pr.Profit.Variance,

			pr.GainRate.Total,
			float64(pr.GainRate.Count),
			pr.GainRate.Average,
			pr.GainRate.Variance,

			pr.LossRate.Total,
			float64(pr.LossRate.Count),
			pr.LossRate.Average,
			pr.LossRate.Variance,

			pr.ProfitRate.Total,
			float64(pr.ProfitRate.Count),
			pr.ProfitRate.Average,
			pr.ProfitRate.Variance,

			// Total
			pr.TotalBuy,
			pr.TotalSell,
			pr.TotalReturnRate,

			// Trade
			pr.WinRate,
			pr.PayoffRatio,
			pr.ProfitFactor,

			// Risk
			pr.SharpRatio,
			pr.SortinoRatio,
			pr.MaxDrawDown,
		}
	}

	return Vector{
		// Total
		pr.TotalBuy,
		pr.TotalSell,
		pr.TotalReturnRate,

		// Trade
		pr.WinRate,
		pr.PayoffRatio,
		pr.ProfitFactor,

		// Risk
		pr.SharpRatio,
		pr.SortinoRatio,
		pr.MaxDrawDown,
	}
}

func safeDivide(condition bool, value1, value2, defaultValue float64) float64 {
	if condition {
		return value1 / value2
	}
	return defaultValue
}

func CreateTradingEvaluationResult(buy Vector, sell Vector) (*TradingEvaluationResult, error) {
	pr := TradingEvaluationResult{}
	if len(buy) != len(sell) {
		return nil, errors.New("buy sell vector length is not matched")
	}

	maxDrawDown := 0.0

	for i := 0; i < len(buy); i++ {
		pr.TotalSell += sell[i]
		pr.TotalBuy += buy[i]
		diff := math.NaN()
		diff_ratio := math.NaN()

		if buy[i] > 0 {
			diff = sell[i] - buy[i]
			diff_ratio = diff / buy[i]
		} else {
			continue
		}

		if diff >= 0 {
			pr.Gain.Add(diff)
			pr.GainRate.Add(diff_ratio)
		} else {
			pr.Loss.Add(diff)
			pr.LossRate.Add(diff_ratio)
			if diff_ratio < maxDrawDown {
				maxDrawDown = diff_ratio
			}
		}
		pr.Profit.Add(diff)
		pr.ProfitRate.Add(diff_ratio)
	}

	pr.Gain.FinalizeAverageCalculation()
	pr.Loss.FinalizeAverageCalculation()
	pr.Profit.FinalizeAverageCalculation()
	pr.GainRate.FinalizeAverageCalculation()
	pr.LossRate.FinalizeAverageCalculation()
	pr.ProfitRate.FinalizeAverageCalculation()
	pr.Gain.Variance = 0
	pr.Loss.Variance = 0
	pr.Profit.Variance = 0
	pr.GainRate.Variance = 0
	pr.LossRate.Variance = 0
	pr.ProfitRate.Variance = 0

	for i := 0; i < len(buy); i++ {
		diff := math.NaN()
		diff_ratio := math.NaN()

		if buy[i] > 0 {
			diff = sell[i] - buy[i]
			diff_ratio = diff / buy[i]
		}

		if math.IsNaN(diff) {
			continue
		}

		if diff >= 0 {
			tmp := diff - pr.Gain.Average
			pr.Gain.Variance += tmp * tmp
			tmp = diff_ratio - pr.GainRate.Average
			pr.GainRate.Variance += tmp * tmp
		} else {
			tmp := diff - pr.Loss.Average
			pr.Loss.Variance += tmp * tmp
			tmp = diff_ratio - pr.LossRate.Average
			pr.LossRate.Variance += tmp * tmp
		}
		tmp := diff - pr.Profit.Average
		pr.Profit.Variance += tmp * tmp
		tmp = diff_ratio - pr.ProfitRate.Average
		pr.ProfitRate.Variance += tmp * tmp
	}

	pr.Gain.FinalizeVarianceCalculation()
	pr.Loss.FinalizeVarianceCalculation()
	pr.Profit.FinalizeVarianceCalculation()
	pr.GainRate.FinalizeVarianceCalculation()
	pr.LossRate.FinalizeVarianceCalculation()
	pr.ProfitRate.FinalizeVarianceCalculation()

	pr.WinRate = safeDivide(pr.Profit.Count != 0, float64(pr.Gain.Count), float64(pr.Profit.Count), math.NaN())
	pr.PayoffRatio = safeDivide(pr.Loss.Average != 0, pr.Gain.Average, -pr.Loss.Average, math.NaN())
	pr.ProfitFactor = safeDivide(pr.Loss.Total != 0, pr.Gain.Total, -pr.Loss.Total, math.NaN())
	pr.TotalReturnRate = safeDivide(pr.TotalBuy != 0, pr.Profit.Total, pr.TotalBuy, 0)
	pr.SharpRatio = safeDivide(pr.ProfitRate.Variance != 0, pr.ProfitRate.Average, pr.ProfitRate.Variance, math.NaN())
	pr.SortinoRatio = safeDivide(pr.LossRate.Variance != 0, pr.ProfitRate.Average, pr.LossRate.Variance, math.NaN())
	pr.MaxDrawDown = maxDrawDown

	return &pr, nil
}

func (pr *TradingEvaluationResult) ToHeader(isDetailOutput bool) []string {
	if isDetailOutput {
		return []string{
			// detail
			"gain.total",
			"gain.count",
			"gain.average",
			"gain.variance",
			"loss.total",
			"loss.count",
			"loss.average",
			"loss.variance",
			"profit.total",
			"profit.count",
			"profit.average",
			"profit.variance",
			"gain_rate.total",
			"gain_rate.count",
			"gain_rate.average",
			"gain_rate.variance",
			"loss_rate.total",
			"loss_rate.count",
			"loss_rate.average",
			"loss_rate.variance",
			"profit_rate.total",
			"profit_rate.count",
			"profit_rate.average",
			"profit_rate.variance",

			// total
			"total_buy",
			"total_sell",
			"total_profit_rate",

			// trade
			"win_rate",
			"payoff_ratio",
			"profit_factor",

			// risk
			"sharp_ratio",
			"sortino_ratio",
			"max_draw_down",
		}
	}

	return []string{
		// total
		"total_buy",
		"total_sell",
		"total_profit_rate",

		// trade
		"win_rate",
		"payoff_ratio",
		"profit_factor",

		// risk
		"sharp_ratio",
		"sortino_ratio",
		"max_draw_down",
	}
}

func (pr *TradingEvaluationResult) PrintTsvHeader(isDetailOutput bool) {
	fmt.Println(strings.Join(pr.ToHeader(isDetailOutput), "\t"))
}

func (pr *TradingEvaluationResult) PrintTsvBody(isDetailOutput bool, precise bool) {
	pr.ToVector(isDetailOutput).PrintTsv(precise)
}
