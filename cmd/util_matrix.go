package cmd

import "math"

type Matrix []Vector

func (data Matrix) InnerProduct() Vector {
	vector := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		vector[i] = 1
		for j := 0; j < len(data); j++ {
			vector[i] *= data[j][i]
		}
	}
	return vector
}

func (data Matrix) GreaterEqual() Vector {
	result := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		v0 := data[0][i]
		v1 := data[1][i]
		if math.IsNaN(v0) || math.IsNaN(v1) {
			result[i] = math.NaN()
		} else if v0 >= v1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func (data Matrix) GreaterThan() Vector {
	result := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		v0 := data[0][i]
		v1 := data[1][i]
		if math.IsNaN(v0) || math.IsNaN(v1) {
			result[i] = math.NaN()
		} else if v0 > v1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}
