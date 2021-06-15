package pkg

import (
	"errors"
	"fmt"
)

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
