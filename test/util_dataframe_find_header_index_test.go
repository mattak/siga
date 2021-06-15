package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"testing"
)

func TestFindHeaderIndex(t *testing.T) {
	df := &dataframe.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
	}

	if df.FindHeaderIndex("date") != 0 {
		t.Fatal("date index not expected")
	}
	if df.FindHeaderIndex("column1") != 1 {
		t.Fatal("column1 index not expected")
	}
	if df.FindHeaderIndex("column2") != 2 {
		t.Fatal("column2 index not expected")
	}
	if df.FindHeaderIndex("column") != -1 {
		t.Fatal("column index not expected")
	}
}
