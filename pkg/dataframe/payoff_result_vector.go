package dataframe

import (
	"math"
)

func PayoffRatio(data Vector) PayoffResult {
	profit := 0.0
	loss := 0.0

	result := PayoffResult{
		PayoffRatio:   math.NaN(),
		ProfitAverage: math.NaN(),
		LossAverage:   math.NaN(),
		ProfitCount:   0,
		LossCount:     0,
	}

	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			profit += data[i]
			result.ProfitCount++
		} else {
			loss += data[i]
			result.LossCount++
		}
	}

	if result.ProfitCount != 0 {
		result.ProfitAverage = profit / float64(result.ProfitCount)
	}
	if result.LossCount != 0 {
		result.LossAverage = loss / float64(result.LossCount)
	}

	if result.LossAverage == 0 || math.IsNaN(result.LossAverage) {
		result.PayoffRatio = math.NaN()
	} else if result.ProfitAverage > 0 {
		result.PayoffRatio = result.ProfitAverage / -result.LossAverage
	}

	return result
}
