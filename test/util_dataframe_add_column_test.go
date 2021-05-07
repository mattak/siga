package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestAddColumn(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1},
		{2},
	}

	err := df.AddColumn("column2", []float64{10, 20})
	if err != nil {
		t.Fatal("wrong column selected, but error was not returned")
	}
	if len(df.Headers) != 3 || df.Headers[2] != "column2" {
		t.Fatal("wrong headers")
	}
	if len(df.Data[0]) != 2 || len(df.Data[1]) != 2 {
		t.Fatal("wrong headers")
	}
	if df.Data[0][1] != 10 || df.Data[1][1] != 20 {
		t.Fatal("data is wrong")
	}
}
