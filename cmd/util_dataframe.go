package cmd

import (
	"errors"
	"fmt"
	"strings"
)

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

func (df *DataFrame) ExtractColumn(columnName string) (Vector, error) {
	columnIndex := df.FindColumnIndex(columnName)
	if columnIndex == -1 {
		return nil, errors.New(fmt.Sprintf("Not found index of header name: %s\n", columnName))
	}

	data := CreateVector(len(df.Labels))
	for i := 0; i < len(df.Labels); i++ {
		data[i] = df.Data[i][columnIndex]
	}
	return data, nil
}

func (df *DataFrame) Clone() *DataFrame {
	newDf := &DataFrame{}
	newDf.Headers = df.Headers
	newDf.Labels = df.Labels
	newDf.Data = df.Data
	return newDf
}

func (df *DataFrame) AddColumn(newHeader string, newDataColumn []float64) error {
	if len(newDataColumn) != len(df.Labels) {
		errors.New("column length is not matched")
	}

	df.Headers = append(df.Headers, newHeader)
	for i := 0; i < len(df.Data); i++ {
		rowLength := len(df.Data[i])
		newRow := make([]float64, rowLength+1)
		for j := 0; j < rowLength; j++ {
			newRow[j] = df.Data[i][j]
		}
		newRow[rowLength] = newDataColumn[i]
		df.Data[i] = newRow
	}

	return nil
}

func (df *DataFrame) SelectColumn(selectionHeaders ...string) error {
	// find column indexes
	selectionIndexes := make([]int, len(selectionHeaders))
	for i := 0; i < len(selectionHeaders); i++ {
		index := df.FindHeaderIndex(selectionHeaders[i])
		if index < 0 {
			return errors.New(fmt.Sprintf("selection index not found: %s", selectionHeaders[i]))
		}
		selectionIndexes[i] = index
	}

	// update data
	for i := 0; i < len(df.Data); i++ {
		rows := make([]float64, len(selectionHeaders))
		for j := 0; j < len(selectionHeaders); j++ {
			index := selectionIndexes[j] - 1
			rows[j] = df.Data[i][index]
		}
		df.Data[i] = rows
	}

	// update headers
	headers := make([]string, len(selectionHeaders)+1)
	headers[0] = df.Headers[0]
	for i := 0; i < len(selectionHeaders); i++ {
		headers[i+1] = selectionHeaders[i]
	}
	df.Headers = headers

	return nil
}

func (df *DataFrame) PrintTsv() {
	fmt.Println(strings.Join(df.Headers, "\t"))
	for i := 0; i < len(df.Labels); i++ {
		fmt.Print(df.Labels[i])
		for j := 0; j < len(df.Data[i]); j++ {
			fmt.Printf("\t%.3f", df.Data[i][j])
		}
		fmt.Println()
	}
}
