package toolkit

func (data Vector) Reverse() {
	j := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		tmp := data[i]
		data[i] = data[j]
		data[j] = tmp
		j--
	}
}

func (data Vector) Fill(value float64) {
	for i := 0; i < len(data); i++ {
		data[i] = value
	}
}
