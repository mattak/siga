package toolkit

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Vector []float64

func CreateVector(size int) Vector {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.NaN()
	}
	return data
}

func CreateVectorWithValue(size int, value float64) Vector {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = value
	}
	return data
}

func CreateVectorWithRandom(size int, fromValue float64, toValue float64) Vector {
	rand.Seed(time.Now().UTC().UnixNano())
	if toValue < fromValue {
		tmp := fromValue
		fromValue = toValue
		toValue = tmp
	}
	span := toValue - fromValue

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Float64()*span + fromValue
	}
	return data
}

func CreateVectorWithRandomWalk(size int, setting RwalkSetting) Vector {
	rand.Seed(time.Now().UTC().UnixNano())
	data := make([]float64, size)

	// u: 15:0.05,0.15
	// d: 15:0,0.05
	// l: 35:-0.05,0
	// r: 35:-0.15,-0.05
	data[0] = setting.StartValue
	for i := 1; i < size; i++ {
		ratio := setting.GetNextRange().CreateRand()
		data[i] = (ratio + 1) * data[i-1]
	}

	return data
}

func (data Vector) Reverse() {
	j := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		tmp := data[i]
		data[i] = data[j]
		data[j] = tmp
		j--
	}
}

func (data Vector) Fill(value float64) {
	for i := 0; i < len(data); i++ {
		data[i] = value
	}
}

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

func (data Vector) SigmaAnomalies(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(data)-span; i++ {
		mean := data.Mean(i+1, span)
		deviation := data.Deviation(i+1, span)
		diff := data[i] - mean

		if deviation != 0 {
			result[i] = diff / deviation
		} else {
			result[i] = 0
		}
	}
	return result
}

func (data Vector) Means(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.Mean(i, span)
	}
	return result
}

func (data Vector) Deviations(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.Deviation(i, span)
	}
	return result
}

func (data Vector) Shift(offset int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(data); i++ {
		p := i - offset
		if p >= 0 && p < len(data) {
			result[i] = data[p]
		} else {
			result[i] = math.NaN()
		}
	}
	return result
}

func (data Vector) Invert() Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(data); i++ {
		if data[i] == 0.0 {
			result[i] = 1.0
		} else if data[i] == 1.0 {
			result[i] = 0.0
		} else {
			result[i] = data[i]
		}
	}
	return result
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

func (price Vector) DollarCostAverage() Vector {
	result := make(Vector, len(price))

	total_invest := 0.0
	sum_volume := 0.0
	for i := 0; i < len(price); i++ {
		if math.IsNaN(price[i]) || price[i] == 0 {
			if i > 0 {
				result[i] = result[i-1]
			} else {
				result[i] = math.NaN()
			}
			continue
		}

		total_invest += 1.0
		volume := 1.0 / price[i]
		sum_volume += volume

		result[i] = total_invest / sum_volume
	}

	return result
}

func (data Vector) PrintTsv(precise bool) {
	startFloatFormat := "%.3f"
	floatFormat := "\t%.3f"

	if precise {
		startFloatFormat = "%f"
		floatFormat = "\t%f"
	}

	if len(data) > 0 {
		fmt.Printf(startFloatFormat, data[0])
	}

	for i := 1; i < len(data); i++ {
		fmt.Printf(floatFormat, data[i])
	}

	if len(data) > 0 {
		fmt.Println()
	}
}
