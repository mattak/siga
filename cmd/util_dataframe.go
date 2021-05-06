package cmd

import "strings"

type DataFrame struct {
	Headers []string    `json:"headers"`
	Labels  []string    `json:"labels"`
	Data    [][]float64 `json:"data"`
}

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

func (df *DataFrame) FindHeaderIndex(name string) int {
	for i, v := range df.Headers {
		if strings.Compare(v, name) == 0 {
			return i
		}
	}
	return -1
}
