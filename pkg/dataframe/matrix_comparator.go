package dataframe

import (
	"github.com/mattak/siga/pkg/util"
	"math"
)

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
		result[i] = util.BoolToFloat64(isAllSame)
	}
	return result
}
