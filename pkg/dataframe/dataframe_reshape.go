package dataframe

func (df *DataFrame) SelectRecords(indexes ...int) {
	labels := make([]string, len(indexes))
	data := make([][]float64, len(indexes))
	for i := 0; i < len(indexes); i++ {
		labels[i] = df.Labels[indexes[i]]
		data[i] = df.Data[indexes[i]]
	}
	df.Labels = labels
	df.Data = data
}

func (df *DataFrame) Reverse() {
	j := len(df.Labels) - 1
	for i := 0; i < len(df.Labels)/2; i++ {
		tmpLabel := df.Labels[i]
		df.Labels[i] = df.Labels[j]
		df.Labels[j] = tmpLabel

		for column := 0; column < len(df.Headers)-1; column++ {
			tmpData := df.Data[i][column]
			df.Data[i][column] = df.Data[j][column]
			df.Data[j][column] = tmpData
		}
		j--
	}
}

func (df *DataFrame) Take(size int) {
	if len(df.Labels) <= size {
		return
	}
	df.Labels = df.Labels[0:size]
	df.Data = df.Data[0:size]
}
