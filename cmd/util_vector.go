package cmd

import (
	"math"
)

type Vector []float64

func CreateVector(size int) Vector {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.NaN()
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
