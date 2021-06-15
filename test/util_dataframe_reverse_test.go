package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"testing"
)

func TestDataFrameReverse(t *testing.T) {
	df := &dataframe.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
	}

	df.Reverse()

	if df.Labels[0] != "seq2" || df.Labels[1] != "seq1" {
		t.Fatal("date index not expected")
	}
	if df.Data[0][0] != 2 || df.Data[1][0] != 1 {
		t.Fatal("column1 data is not expected")
	}
	if df.Data[0][1] != 20 || df.Data[1][1] != 10 {
		t.Fatal("column2 data is not expected")
	}
}
