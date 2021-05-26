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

func (data Matrix) Divide() Vector {
	vector := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		vector[i] = data[0][i]

		if math.IsNaN(vector[i]) {
			continue
		}

		for j := 1; j < len(data); j++ {
			if math.IsNaN(data[j][i]) {
				vector[i] = math.NaN()
			} else if data[j][i] == 0 {
				vector[i] = math.NaN()
			} else {
				vector[i] /= data[j][i]
			}
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

func (data Matrix) LessEqual() Vector {
	result := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		v0 := data[0][i]
		v1 := data[1][i]
		if math.IsNaN(v0) || math.IsNaN(v1) {
			result[i] = math.NaN()
		} else if v0 <= v1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func (data Matrix) LessThan() Vector {
	result := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		v0 := data[0][i]
		v1 := data[1][i]
		if math.IsNaN(v0) || math.IsNaN(v1) {
			result[i] = math.NaN()
		} else if v0 < v1 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func (data Matrix) Equal() Vector {
	result := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		v0 := data[0][i]
		isAllSame := true

		if math.IsNaN(v0) {
			for j := 1; j < len(data); j++ {
				if !math.IsNaN(data[j][i]) {
					isAllSame = false
					break
				}
			}
		} else {
			for j := 1; j < len(data); j++ {
				if v0 != data[j][i] {
					isAllSame = false
					break
				}
			}
		}
		result[i] = BoolToFloat64(isAllSame)
	}
	return result
}

func (data Matrix) FilterIndex() []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		if math.IsNaN(data[0][i]) && math.IsNaN(data[1][i]) || data[0][i] == data[1][i] {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (data Matrix) FilterNotIndex() []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		if !(math.IsNaN(data[0][i]) && math.IsNaN(data[1][i]) || data[0][i] == data[1][i]) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
