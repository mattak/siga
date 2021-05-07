package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestSelectSingleColumn(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
	}

	err := df.SelectColumn("column2")
	if err != nil {
		t.Fatal(err)
	}

	if len(df.Headers) != 2 || df.Headers[0] != "date" || df.Headers[1] != "column2" {
		t.Fatal("header length is not expected")
	}
	if len(df.Labels) != 2 {
		t.Fatal("labels length is not expected")
	}
	if len(df.Data) != 2 || len(df.Data[0]) != 1 || len(df.Data[1]) != 1 {
		t.Fatal("data length is not expected")
	}
	if df.Data[0][0] != 10.0 {
		t.Fatal("data[0][0] is not expected")
	}
	if df.Data[1][0] != 20.0 {
		t.Fatal("data[1][0] is not expected")
	}
}

func TestSelectShuffleColumn(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
	}

	err := df.SelectColumn("column2", "column1")
	if err != nil {
		t.Fatal(err)
	}

	if len(df.Headers) != 3 || df.Headers[0] != "date" || df.Headers[1] != "column2" || df.Headers[2] != "column1" {
		t.Fatal("header length is not expected")
	}
	if len(df.Labels) != 2 {
		t.Fatal("labels length is not expected")
	}
	if len(df.Data) != 2 || len(df.Data[0]) != 2 || len(df.Data[1]) != 2 {
		t.Fatal("data length is not expected")
	}
	if df.Data[0][0] != 10.0 {
		t.Fatal("data[0][0] is not expected")
	}
	if df.Data[0][1] != 1.0 {
		t.Fatal("data[0][1] is not expected")
	}
	if df.Data[1][0] != 20.0 {
		t.Fatal("data[1][0] is not expected")
	}
	if df.Data[1][1] != 2.0 {
		t.Fatal("data[1][1] is not expected")
	}
}

func TestSelectWrongColumn(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
	}

	err := df.SelectColumn("column")
	if err == nil {
		t.Fatal("wrong column selected, but error was not returned")
	}
}
