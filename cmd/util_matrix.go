package cmd

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
