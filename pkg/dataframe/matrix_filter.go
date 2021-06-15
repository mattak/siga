package dataframe

import "math"

func (data Matrix) FilterIndexByAnd(targetValue float64) []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		ok := true
		for c := 0; c < len(data); c++ {
			if !(math.IsNaN(targetValue) && math.IsNaN(data[c][i]) || targetValue == data[c][i]) {
				ok = false
				break
			}
		}
		if ok {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (data Matrix) FilterIndexByOr(targetValue float64) []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		ok := false
		for c := 0; c < len(data); c++ {
			if math.IsNaN(targetValue) && math.IsNaN(data[c][i]) || targetValue == data[c][i] {
				ok = true
				break
			}
		}
		if ok {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (data Matrix) FilterIndexByNotAnd(targetValue float64) []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		ok := true
		for c := 0; c < len(data); c++ {
			if !(math.IsNaN(targetValue) && math.IsNaN(data[c][i]) || targetValue == data[c][i]) {
				ok = false
				break
			}
		}
		if !ok {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (data Matrix) FilterIndexByNotOr(targetValue float64) []int {
	indexes := []int{}
	for i := 0; i < len(data[0]); i++ {
		ok := false
		for c := 0; c < len(data); c++ {
			if math.IsNaN(targetValue) && math.IsNaN(data[c][i]) || targetValue == data[c][i] {
				ok = true
				break
			}
		}
		if !ok {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
