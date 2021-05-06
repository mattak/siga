package cmd

import "strconv"

func ParseInt(text string) int {
	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func ParseInt64(text string) int64 {
	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func ParseFloat64(text string) float64 {
	n, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(err)
	}
	return n
}
