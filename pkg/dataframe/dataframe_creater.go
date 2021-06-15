package dataframe

import (
	"strconv"
)

func CreateDataFrame(headers []string, data []Vector) DataFrame {
	df := DataFrame{}
	df.Headers = headers

	df.Labels = make([]string, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		df.Labels[i] = strconv.Itoa(i)
	}

	df.Data = make([][]float64, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		df.Data[i] = make([]float64, len(data))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			df.Data[j][i] = data[i][j]
		}
	}

	return df
}
