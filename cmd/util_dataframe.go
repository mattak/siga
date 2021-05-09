package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DataFrame struct {
	Headers []string    `json:"headers"`
	Labels  []string    `json:"labels"`
	Data    [][]float64 `json:"data"`
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

func (df *DataFrame) ExtractMatrixByColumnNameOrValue(args []string) Matrix {
	matrix := make(Matrix, len(args))
	for i := 0; i < len(args); i++ {
		matrix[i] = CreateVector(len(df.Labels))

		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			matrix[i] = vector
		} else {
			v, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
			} else {
				matrix[i].Fill(v)
			}
		}
	}
	return matrix
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
