package pkg

func (df *DataFrame) FindHeaderIndex(name string) int {
	for i, v := range df.Headers {
		if v == name {
			return i
		}
	}
	return -1
}

func (df *DataFrame) FindColumnIndex(name string) int {
	for i, v := range df.Headers {
		if v == name {
			return i - 1
		}
	}
	return -1
}

func (df *DataFrame) FindLabelIndex(name string) int {
	for i, v := range df.Labels {
		if v == name {
			return i
		}
	}
	return -1
}
