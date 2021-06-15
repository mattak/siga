package dataframe

import "math"

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

func (vector Vector) NormalizeByStart() Vector {
	result := make(Vector, len(vector))

	for i := 0; i < len(result); i++ {
		result[i] = vector[i] / vector[0]
	}

	return result
}
