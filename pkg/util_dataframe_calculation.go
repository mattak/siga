package pkg

func (df *DataFrame) Sum(column, from, size int) float64 {
	sum := 0.0
	for i := from; i < from+size; i++ {
		sum += df.Data[i][column]
	}
	return sum
}

func (df *DataFrame) Mean(column, from, size int) float64 {
	return df.Sum(column, from, size) / float64(size)
}