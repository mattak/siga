package dataframe

import (
	"errors"
	"strconv"
)

func CreateDataFrame(headers []string, data []Vector) (*DataFrame, error) {
	if len(headers) != len(data) {
		return nil, errors.New("DataFrame Header and Data length is not matched")
	}

	df := DataFrame{}
	df.Headers = append([]string{"index"}, headers...)

	df.Labels = make([]string, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		df.Labels[i] = strconv.Itoa(i + 1)
	}

	df.Data = make([][]float64, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		df.Data[i] = make([]float64, len(data))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			df.Data[j][i] = data[i][j]
		}
	}

	return &df, nil
}

func CreateDataFrameWithLabels(headers []string, labels []string, data []Vector) (*DataFrame, error) {
	if len(headers)-1 != len(data) {
		return nil, errors.New("headers and data length is not matched")
	}
	if len(data) > 0 && len(labels) != len(data[0]) {
		return nil, errors.New("labels and data length is not matched")
	}

	df := DataFrame{}
	df.Headers = headers
	df.Labels = labels

	df.Data = make([][]float64, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		df.Data[i] = make([]float64, len(data))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			df.Data[j][i] = data[i][j]
		}
	}

	return &df, nil
}

func (frame *DataFrame) Clone() *DataFrame {
	headers := make([]string, len(frame.Headers))
	labels := make([]string, len(frame.Labels))
	data := make([][]float64, len(frame.Data))

	for i := 0; i< len(frame.Headers); i++ {
		headers[i] = frame.Headers[i]
	}
	for i := 0; i< len(frame.Labels); i++ {
		labels[i] = frame.Labels[i]
	}
	for i := 0; i< len(frame.Data); i++ {
		data[i] = make([]float64, len(frame.Data[i]))
		for j := 0; j < len(frame.Data[i]); j++ {
			data[i][j] = frame.Data[i][j];
		}
	}

	return &DataFrame{
		Headers: headers,
		Labels: labels,
		Data: data,
	}
}
