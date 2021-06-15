package dataframe

func (data Vector) SimpleMovingAverage(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.Mean(i, span)
	}
	return result
}

func (data Vector) HarmonicMovingAverage(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.HarmonicMean(i, span)
	}
	return result
}

func (data Vector) ValueAverage(span int) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.ValueMean(i, span)
	}
	return result
}

func (data Vector) MartinegaleAverage(span int, threshold float64) Vector {
	result := CreateVector(len(data))
	for i := 0; i < len(result)-span+1; i++ {
		result[i] = data.MartinegaleMean(i, span, threshold)
	}
	return result
}
