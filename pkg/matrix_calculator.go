package pkg

import "math"

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

func (data Matrix) Add() Vector {
	vector := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		vector[i] = data[0][i]

		if math.IsNaN(vector[i]) {
			continue
		}

		for j := 1; j < len(data); j++ {
			if math.IsNaN(data[j][i]) {
				vector[i] = math.NaN()
				break
			}

			vector[i] += data[j][i]
		}
	}

	return vector
}

func (data Matrix) Subtract() Vector {
	vector := CreateVector(len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		vector[i] = data[0][i]

		if math.IsNaN(vector[i]) {
			continue
		}

		for j := 1; j < len(data); j++ {
			if math.IsNaN(data[j][i]) {
				vector[i] = math.NaN()
				break
			}

			vector[i] -= data[j][i]
		}
	}

	return vector
}
