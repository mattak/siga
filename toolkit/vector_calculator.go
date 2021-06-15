package toolkit

import (
	"math"
)

func (data Vector) HasAnyValue(value float64) bool {
	if math.IsNaN(value) {
		for i := 0; i < len(data); i++ {
			if math.IsNaN(data[i]) {
				return true
			}
		}
	} else {
		for i := 0; i < len(data); i++ {
			if data[i] == value {
				return true
			}
		}
	}
	return false
}

func (data Vector) IsAllValue(value float64) bool {
	if math.IsNaN(value) {

		for i := 0; i < len(data); i++ {
			if !math.IsNaN(data[i]) {
				return false
			}
		}
	} else {
		for i := 0; i < len(data); i++ {
			if data[i] != value {
				return false
			}
		}
	}
	return true
}

func (data Vector) Sum(start, length int) float64 {
	result := 0.0
	for i := start; i < start+length; i++ {
		result += data[i]
	}
	return result
}

func (data Vector) Mean(start, length int) float64 {
	return data.Sum(start, length) / float64(length)
}

func (data Vector) HarmonicMean(start, length int) float64 {
	// D = a1, a2, a3
	// D' = 1/a1, 1/a2, 1/a3
	// R = Count(D')/Sum(D')
	sum := 0.0
	for i := start; i < start+length; i++ {
		sum += 1 / data[i]
	}
	if sum == 0.0 {
		return math.NaN()
	}
	return float64(length) / sum
}

func (data Vector) ValueMean(start, length int) float64 {
	target_value := 0.0
	amounts := 0.0
	invests := 0.0

	for i := start; i < start+length; i++ {
		target_value += 1.0
		current_value := amounts * data[i]
		diff := target_value - current_value
		if diff > 0 {
			amount := diff / data[i]
			amounts += amount
			invests += diff
		}
	}

	if amounts > 0 {
		return invests / amounts
	}
	return math.NaN()
}

func (data Vector) MartinegaleMean(start, length int, threshold float64) float64 {
	if length < 1 || len(data) < 1 {
		return math.NaN()
	}

	last_value := data[0] * threshold
	step := 1.0
	invests := step
	amounts := step / data[0]

	for i := start + 1; i < start+length; i++ {
		if data[i] < last_value {
			step *= 2.0
			invests += step
			amounts += step / data[i]
			last_value = (invests / amounts) * threshold
		}
	}

	if amounts > 0 {
		return invests / amounts
	}

	return math.NaN()
}

func (data Vector) DeviationSquare(start, length int) float64 {
	mean := data.Mean(start, length)
	result := 0.0

	for i := start; i < start+length; i++ {
		diff := data[i] - mean
		result += diff * diff
	}

	return result / float64(length)
}

func (data Vector) Deviation(start, length int) float64 {
	return math.Sqrt(data.DeviationSquare(start, length))
}

func (data Vector) ProfitFactor() float64 {
	positive := 0.0
	negative := 0.0

	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			positive += data[i]
		} else {
			negative += data[i]
		}
	}

	if negative == 0 {
		return math.Inf(1)
	}

	return positive / -negative
}

func (data Vector) PayoffRatio() PayoffResult {
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
